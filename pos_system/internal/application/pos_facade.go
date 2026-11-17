package application

import (
	"fmt"
	"pos_system/internal/domain"
)

type POSFacade struct {
	studentRepo   domain.StudentRepository
	criterionRepo domain.CriterionRepository
	evalRepo      domain.EvaluationRepository
	summarySvc    *domain.SummaryService
}

func NewPOSFacade(sr domain.StudentRepository, cr domain.CriterionRepository, er domain.EvaluationRepository, ss *domain.SummaryService) *POSFacade {
	return &POSFacade{
		studentRepo:   sr,
		criterionRepo: cr,
		evalRepo:      er,
		summarySvc:    ss,
	}
}

func (f *POSFacade) AddStudent(id, name string) error {
	student := domain.NewStudent(id, name)
	return f.studentRepo.Save(student)
}

func (f *POSFacade) AddCriterion(id, name string, maxScore, weight float64) error {
	crit, err := domain.NewCriterion(id, name, maxScore, weight)
	if err != nil {
		return err
	}
	evals, err := f.evalRepo.FindAll()
	if err == nil {
		for _, e := range evals {
			if e.CriterionID() == id && e.ScoreValue() > maxScore {
				return fmt.Errorf("конфликт: оценка студента %s (%.2f) превышает новый максимум %.2f", e.StudentID(), e.ScoreValue(), maxScore)
			}
		}
	}
	return f.criterionRepo.Save(crit)
}
func (f *POSFacade) EvaluateStudent(studentID, criterionID string, score float64) error {
	_, err := f.studentRepo.FindByID(studentID)
	if err != nil {
		return fmt.Errorf("студент не найден: %v", err)
	}
	crit, err := f.criterionRepo.FindByID(criterionID)
	if err != nil {
		return fmt.Errorf("критерий не найден: %v", err)
	}
	eval, err := domain.NewEvaluation(studentID, crit, score)
	if err != nil {
		return err
	}
	return f.evalRepo.Save(eval)
}

func (f *POSFacade) GetStudentSummary(studentID string) (float64, error) {
	_, err := f.studentRepo.FindByID(studentID)
	if err != nil {
		return 0, fmt.Errorf("студент не найден: %v", err)
	}
	evals, err := f.evalRepo.FindByStudentID(studentID)
	if err != nil {
		return 0, err
	}
	allCriteria, err := f.criterionRepo.FindAll()
	if err != nil {
		return 0, err
	}
	critMap := make(map[string]*domain.Criterion)
	for _, c := range allCriteria {
		critMap[c.ID] = c
	}
	total := f.summarySvc.CalculateStudentSummary(studentID, evals, critMap)
	return total, nil
}

func (f *POSFacade) GetAllStudents() ([]StudentDTO, error) {
	students, err := f.studentRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return ToStudentDTOList(students), nil
}

func (f *POSFacade) GetStudentEvaluations(studentID string) ([]EvaluationDTO, error) {
	evals, err := f.evalRepo.FindByStudentID(studentID)
	if err != nil {
		return nil, err
	}
	return ToEvaluationDTOList(evals), nil
}

func (f *POSFacade) RemoveEvaluation(studentID, criterionID string) error {
	return f.evalRepo.Delete(studentID, criterionID)
}

func (f *POSFacade) GetAllCriteria() ([]domain.Criterion, error) {
	criteria, err := f.criterionRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var res []domain.Criterion
	for _, c := range criteria {
		res = append(res, *c)
	}
	return res, nil
}
