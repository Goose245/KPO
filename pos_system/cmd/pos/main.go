package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"pos_system/internal/application"
	"pos_system/internal/domain"
	"pos_system/internal/infrastructure/repository"
	"strconv"
	"strings"
)

func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	if err := os.MkdirAll("data", 0755); err != nil {
		fmt.Println("[-] Ошибка создания папки data:", err)
		os.Exit(1)
	}
	studentRepo := repository.NewJSONStudentRepo("data/students.json")
	criterionRepo := repository.NewJSONCriterionRepo("data/criteria.json")
	evalRepo := repository.NewJSONEvaluationRepo("data/evaluation.json")
	strategy := &domain.WeightedAverageStrategy{}
	summaryService := domain.NewSummaryService(strategy)
	app := application.NewPOSFacade(studentRepo, criterionRepo, evalRepo, summaryService)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Система ПОС ---")
		fmt.Println("1. Добавить студента")
		fmt.Println("2. Добавить критерий")
		fmt.Println("3. Выставить или изменить оценку")
		fmt.Println("4. Удалить оценку")
		fmt.Println("5. Посмотреть студентов")
		fmt.Println("6. Посмотреть критерии")
		fmt.Println("7. Посмотреть оценки студента и итог")
		fmt.Println("8. Экспорт итоговой ведомости")
		fmt.Println("0. Выход")
		choice := readInput(reader, "Выберите действие: ")
		switch choice {
		case "1":
			id := readInput(reader, "ID студента: ")
			name := readInput(reader, "Имя студента: ")
			err := app.AddStudent(id, name)
			if err != nil {
				fmt.Println("[-] Ошибка:", err)
			} else {
				fmt.Println("[+] Студент добавлен")
			}
		case "2":
			id := readInput(reader, "ID критерия: ")
			name := readInput(reader, "Название критерия: ")
			maxStr := readInput(reader, "Максимальный балл: ")
			weightStr := readInput(reader, "Вес критерия (например, 0.3): ")
			maxScore, err1 := strconv.ParseFloat(maxStr, 64)
			weight, err2 := strconv.ParseFloat(weightStr, 64)
			if err1 != nil || err2 != nil {
				fmt.Println("[-] Некорректный ввод чисел")
				continue
			}
			err := app.AddCriterion(id, name, maxScore, weight)
			if err != nil {
				fmt.Println("[-] Ошибка:", err)
			} else {
				fmt.Println("[+] Критерий добавлен")
			}
		case "3":
			sid := readInput(reader, "ID студента: ")
			cid := readInput(reader, "ID критерия: ")
			scoreStr := readInput(reader, "Оценка: ")
			score, err := strconv.ParseFloat(scoreStr, 64)
			if err != nil {
				fmt.Println("[-] Некорректный ввод оценки")
				continue
			}
			err = app.EvaluateStudent(sid, cid, score)
			if err != nil {
				fmt.Println("[-] Ошибка:", err)
			} else {
				fmt.Println("[+] Оценка сохранена")
			}
		case "4":
			sid := readInput(reader, "ID студента: ")
			cid := readInput(reader, "ID критерия: ")
			err := app.RemoveEvaluation(sid, cid)
			if err != nil {
				fmt.Println("[-] Ошибка:", err)
			} else {
				fmt.Println("[+] Оценка удалена")
			}
		case "5":
			students, err := app.GetAllStudents()
			if err != nil {
				fmt.Println("[-] Ошибка:", err)
				continue
			}
			fmt.Println("--- Список студентов ---")
			for _, s := range students {
				fmt.Printf("ID: %s, Имя: %s\n", s.ID, s.Name)
			}
		case "6":
			criteria, err := app.GetAllCriteria()
			if err != nil {
				fmt.Println("[-] Ошибка:", err)
				continue
			}
			fmt.Println("--- Рубрикатор ---")
			for _, c := range criteria {
				fmt.Printf("ID: %s, Название: %s, Макс. балл: %.2f, Вес: %.2f\n", c.ID, c.Name, c.MaxScore, c.Weight)
			}
		case "7":
			sid := readInput(reader, "ID студента: ")
			evals, err := app.GetStudentEvaluations(sid)
			if err != nil {
				fmt.Println("[-] Ошибка:", err)
				continue
			}
			fmt.Println("Оценки:")
			for _, e := range evals {
				fmt.Printf("Критерий: %s, Балл: %.2f\n", e.CriterionID, e.Score)
			}
			total, err := app.GetStudentSummary(sid)
			if err != nil {
				fmt.Println("[-] Ошибка расчета итога:", err)
			} else {
				fmt.Printf("Итоговый балл: %.2f\n", total)
			}
		case "8":
			exportPath := "data/summary_export.json"
			students, err := app.GetAllStudents()
			if err != nil {
				fmt.Println("[-] Ошибка получения списка студентов:", err)
				continue
			}
			var exportData []application.SummaryDTO
			for _, s := range students {
				total, err := app.GetStudentSummary(s.ID)
				if err != nil {
					fmt.Printf("[-] Пропуск студента %s: %v\n", s.ID, err)
					continue
				}
				exportData = append(exportData, application.SummaryDTO{
					StudentID:   s.ID,
					StudentName: s.Name,
					TotalScore:  total,
				})
			}
			file, err := os.Create(exportPath)
			if err != nil {
				fmt.Println("[-] Ошибка создания файла экспорта:", err)
				continue
			}
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(exportData); err != nil {
				fmt.Println("[-] Ошибка записи JSON:", err)
			} else {
				fmt.Println("[+] Итоговая ведомость экспортирована в", exportPath)
			}
			file.Close()
		case "0":
			fmt.Println("Выход.")
			return
		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
