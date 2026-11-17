package domain_test

import (
	"math"
	"pos_system/internal/domain"
	"testing"
)

func TestNewEvaluation_Success(t *testing.T) {
	crit, _ := domain.NewCriterion("c1", "Лекции", 10.0, 1.0)
	eval, err := domain.NewEvaluation("s1", crit, 8.5)

	if err != nil {
		t.Fatalf("Ожидался успешный результат, получена ошибка: %v", err)
	}
	if eval.ScoreValue() != 8.5 {
		t.Errorf("Ожидалась оценка 8.5, получено %v", eval.ScoreValue())
	}
}

func TestNewEvaluation_FailsWhenScoreExceedsMax(t *testing.T) {
	crit, _ := domain.NewCriterion("c1", "Лекции", 10.0, 1.0)
	_, err := domain.NewEvaluation("s1", crit, 12.0)
	if err == nil {
		t.Error("Ожидалась ошибка (срабатывание инварианта), но оценка 12/10 была создана")
	}
}

func TestNewEvaluation_FailsWhenScoreIsNegative(t *testing.T) {
	crit, _ := domain.NewCriterion("c1", "Лекции", 10.0, 1.0)
	_, err := domain.NewEvaluation("s1", crit, -2.0)
	if err == nil {
		t.Error("Ожидалась ошибка (срабатывание инварианта), но отрицательная оценка была создана")
	}
}

func TestNewEvaluation_FailsOnNaNAndInf(t *testing.T) {
	crit, _ := domain.NewCriterion("c1", "C1", 10.0, 1.0)
	_, err1 := domain.NewEvaluation("s1", crit, math.NaN())
	_, err2 := domain.NewEvaluation("s1", crit, math.Inf(1))

	if err1 == nil || err2 == nil {
		t.Error("Ожидалась ошибка при передаче NaN или Inf")
	}
}

func TestSummaryService_CalculateStudentSummary(t *testing.T) {
	strategy := &domain.WeightedAverageStrategy{}
	svc := domain.NewSummaryService(strategy)

	c1, _ := domain.NewCriterion("c1", "C1", 10.0, 0.4)
	c2, _ := domain.NewCriterion("c2", "C2", 20.0, 0.6)
	critMap := map[string]*domain.Criterion{"c1": c1, "c2": c2}

	e1, _ := domain.NewEvaluation("s1", c1, 8.0)
	e2, _ := domain.NewEvaluation("s1", c2, 15.0)

	total := svc.CalculateStudentSummary("s1", []domain.Evaluation{*e1, *e2}, critMap)
	expected := 0.77

	if math.Abs(total-expected) > 0.0001 {
		t.Errorf("Ожидалось %.2f, получено %.2f", expected, total)
	}
}

type MockStrategy struct{}

func (m *MockStrategy) Calculate(evals []domain.Evaluation, criteria map[string]*domain.Criterion) float64 {
	return 999.0
}

func TestSummaryService_ChangeStrategy(t *testing.T) {
	svc := domain.NewSummaryService(&MockStrategy{})
	total := svc.CalculateStudentSummary("s1", nil, nil)
	if total != 999.0 {
		t.Errorf("Ожидалось 999.0 (от мок-стратегии), получено %.2f", total)
	}
}
