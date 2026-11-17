package application

import "pos_system/internal/domain"

type StudentDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SummaryDTO struct {
	StudentID   string  `json:"student_id"`
	StudentName string  `json:"student_name"`
	TotalScore  float64 `json:"total_score"`
}

type EvaluationDTO struct {
	StudentID   string  `json:"student_id"`
	CriterionID string  `json:"criterion_id"`
	Score       float64 `json:"score"`
}

func ToStudentDTOList(students []*domain.Student) []StudentDTO {
	result := make([]StudentDTO, len(students))
	for i, s := range students {
		result[i] = StudentDTO{
			ID:   s.ID,
			Name: s.Name,
		}
	}
	return result
}

func ToEvaluationDTOList(evals []domain.Evaluation) []EvaluationDTO {
	result := make([]EvaluationDTO, len(evals))
	for i, e := range evals {
		result[i] = EvaluationDTO{
			StudentID:   e.StudentID(),
			CriterionID: e.CriterionID(),
			Score:       float64(e.ScoreValue()),
		}
	}
	return result
}
