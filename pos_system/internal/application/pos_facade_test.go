package application_test

import (
	"errors"
	"pos_system/internal/application"
	"pos_system/internal/domain"
	"testing"
)

type MemStudentRepo struct{ data map[string]*domain.Student }

func (m *MemStudentRepo) Save(s *domain.Student) error { m.data[s.ID] = s; return nil }
func (m *MemStudentRepo) FindByID(id string) (*domain.Student, error) {
	if s, ok := m.data[id]; ok {
		return s, nil
	}
	return nil, errors.New("not found")
}
func (m *MemStudentRepo) FindAll() ([]*domain.Student, error) { return nil, nil }

type MemCritRepo struct{ data map[string]*domain.Criterion }

func (m *MemCritRepo) Save(c *domain.Criterion) error { m.data[c.ID] = c; return nil }
func (m *MemCritRepo) FindByID(id string) (*domain.Criterion, error) {
	if c, ok := m.data[id]; ok {
		return c, nil
	}
	return nil, errors.New("not found")
}
func (m *MemCritRepo) FindAll() ([]*domain.Criterion, error) { return nil, nil }

type MemEvalRepo struct{ data []domain.Evaluation }

func (m *MemEvalRepo) Save(eval *domain.Evaluation) error {
	for i, e := range m.data {
		if e.StudentID() == eval.StudentID() && e.CriterionID() == eval.CriterionID() {
			m.data[i] = *eval
			return nil
		}
	}
	m.data = append(m.data, *eval)
	return nil
}
func (m *MemEvalRepo) FindByStudentID(id string) ([]domain.Evaluation, error) {
	var res []domain.Evaluation
	for _, e := range m.data {
		if e.StudentID() == id {
			res = append(res, e)
		}
	}
	return res, nil
}
func (m *MemEvalRepo) FindAll() ([]domain.Evaluation, error) { return m.data, nil }
func (m *MemEvalRepo) Delete(sid, cid string) error          { return nil }
func setupFacade() (*application.POSFacade, *MemEvalRepo) {
	evalRepo := &MemEvalRepo{data: make([]domain.Evaluation, 0)}
	studentRepo := &MemStudentRepo{data: map[string]*domain.Student{"s1": domain.NewStudent("s1", "Иван")}}
	c1, _ := domain.NewCriterion("c1", "Критерий 1", 10.0, 1.0)
	critRepo := &MemCritRepo{data: map[string]*domain.Criterion{"c1": c1}}

	f := application.NewPOSFacade(studentRepo, critRepo, evalRepo, domain.NewSummaryService(&domain.WeightedAverageStrategy{}))
	return f, evalRepo
}

func TestFacade_EvaluateUnknownStudentOrCriterion(t *testing.T) {
	f, _ := setupFacade()

	err1 := f.EvaluateStudent("unknown_student", "c1", 5.0)
	if err1 == nil {
		t.Error("Ожидалась ошибка при несуществующем студенте")
	}

	err2 := f.EvaluateStudent("s1", "unknown_crit", 5.0)
	if err2 == nil {
		t.Error("Ожидалась ошибка при несуществующем критерии")
	}
}

func TestFacade_AddCriterion_ProtectsOldEvaluations(t *testing.T) {
	f, _ := setupFacade()
	f.EvaluateStudent("s1", "c1", 9.0)
	err := f.AddCriterion("c1", "Критерий 1", 8.0, 1.0)
	if err == nil {
		t.Error("Ожидалась ошибка: нельзя уменьшить MaxScore ниже уже существующих оценок")
	}
}

func TestFacade_DuplicateEvaluationReplaced(t *testing.T) {
	f, evalRepo := setupFacade()

	f.EvaluateStudent("s1", "c1", 5.0)
	f.EvaluateStudent("s1", "c1", 8.0)

	if len(evalRepo.data) != 1 {
		t.Fatalf("Ожидалась 1 оценка в репозитории, сохранено %d", len(evalRepo.data))
	}
	if evalRepo.data[0].ScoreValue() != 8.0 {
		t.Errorf("Ожидался исправленный балл 8.0, получено %v", evalRepo.data[0].ScoreValue())
	}
}
