package domain

type CalculationStrategy interface {
	Calculate(evaluations []Evaluation, criteria map[string]*Criterion) float64
}

type WeightedAverageStrategy struct{}

func (s *WeightedAverageStrategy) Calculate(evaluations []Evaluation, criteria map[string]*Criterion) float64 {
	var total float64
	for _, eval := range evaluations {
		if crit, exists := criteria[eval.CriterionID()]; exists {
			normalized := float64(eval.ScoreValue()) / crit.MaxScore
			total += normalized * crit.Weight
		}
	}
	return total
}

type SummaryService struct {
	strategy CalculationStrategy
}

func NewSummaryService(strategy CalculationStrategy) *SummaryService {
	return &SummaryService{strategy: strategy}
}

func (ss *SummaryService) CalculateStudentSummary(studentID string, evals []Evaluation, criteria map[string]*Criterion) float64 {
	var studentEvals []Evaluation
	for _, e := range evals {
		if e.StudentID() == studentID {
			studentEvals = append(studentEvals, e)
		}
	}
	return ss.strategy.Calculate(studentEvals, criteria)
}
