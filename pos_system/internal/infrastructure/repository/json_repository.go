package repository

import (
	"encoding/json"
	"errors"
	"os"
	"pos_system/internal/domain"
	"sync"
)

func loadFromFile(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, v)
}

func saveToFile(filename string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

type JSONStudentRepo struct {
	filename string
	mu       sync.RWMutex
}

func NewJSONStudentRepo(filename string) *JSONStudentRepo {
	return &JSONStudentRepo{filename: filename}
}

func (r *JSONStudentRepo) Save(student *domain.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var students []*domain.Student
	if err := loadFromFile(r.filename, &students); err != nil {
		return err
	}
	updated := false
	for i, s := range students {
		if s.ID == student.ID {
			students[i] = student
			updated = true
			break
		}
	}
	if !updated {
		students = append(students, student)
	}
	return saveToFile(r.filename, students)
}

func (r *JSONStudentRepo) FindByID(id string) (*domain.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var students []*domain.Student
	if err := loadFromFile(r.filename, &students); err != nil {
		return nil, err
	}
	for _, s := range students {
		if s.ID == id {
			return s, nil
		}
	}
	return nil, errors.New("студент не найден")
}

func (r *JSONStudentRepo) FindAll() ([]*domain.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var students []*domain.Student
	if err := loadFromFile(r.filename, &students); err != nil {
		return nil, err
	}
	return students, nil
}

type JSONCriterionRepo struct {
	filename string
	mu       sync.RWMutex
}

func NewJSONCriterionRepo(filename string) *JSONCriterionRepo {
	return &JSONCriterionRepo{filename: filename}
}

func (r *JSONCriterionRepo) Save(criterion *domain.Criterion) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var criteria []*domain.Criterion
	if err := loadFromFile(r.filename, &criteria); err != nil {
		return err
	}
	updated := false
	for i, c := range criteria {
		if c.ID == criterion.ID {
			criteria[i] = criterion
			updated = true
			break
		}
	}
	if !updated {
		criteria = append(criteria, criterion)
	}

	return saveToFile(r.filename, criteria)
}

func (r *JSONCriterionRepo) FindByID(id string) (*domain.Criterion, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var criteria []*domain.Criterion
	if err := loadFromFile(r.filename, &criteria); err != nil {
		return nil, err
	}
	for _, c := range criteria {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, errors.New("критерий не найден")
}

func (r *JSONCriterionRepo) FindAll() ([]*domain.Criterion, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var criteria []*domain.Criterion
	if err := loadFromFile(r.filename, &criteria); err != nil {
		return nil, err
	}
	return criteria, nil
}

type JSONEvaluationRepo struct {
	filename string
	mu       sync.RWMutex
}

func NewJSONEvaluationRepo(filename string) *JSONEvaluationRepo {
	return &JSONEvaluationRepo{filename: filename}
}

func (r *JSONEvaluationRepo) Save(eval *domain.Evaluation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var evals []domain.Evaluation
	if err := loadFromFile(r.filename, &evals); err != nil {
		return err
	}

	updated := false
	for i, e := range evals {
		if e.StudentID() == eval.StudentID() && e.CriterionID() == eval.CriterionID() {
			evals[i] = *eval
			updated = true
			break
		}
	}
	if !updated {
		evals = append(evals, *eval)
	}

	return saveToFile(r.filename, evals)
}

func (r *JSONEvaluationRepo) FindByStudentID(studentID string) ([]domain.Evaluation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var allEvals []domain.Evaluation
	if err := loadFromFile(r.filename, &allEvals); err != nil {
		return nil, err
	}
	var result []domain.Evaluation
	for _, e := range allEvals {
		if e.StudentID() == studentID {
			result = append(result, e)
		}
	}
	return result, nil
}

func (r *JSONEvaluationRepo) FindAll() ([]domain.Evaluation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var allEvals []domain.Evaluation
	if err := loadFromFile(r.filename, &allEvals); err != nil {
		return nil, err
	}
	return allEvals, nil
}

func (r *JSONEvaluationRepo) Delete(studentID, criterionID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var evals []domain.Evaluation
	if err := loadFromFile(r.filename, &evals); err != nil {
		return err
	}
	var newEvals []domain.Evaluation
	for _, e := range evals {
		if e.StudentID() == studentID && e.CriterionID() == criterionID {
			continue
		}
		newEvals = append(newEvals, e)
	}
	return saveToFile(r.filename, newEvals)
}
