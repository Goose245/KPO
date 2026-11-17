package domain

import (
	"encoding/json"
	"fmt"
	"math"
)

type Score float64

type Evaluation struct {
	studentID   string
	criterionID string
	score       Score
}

func NewEvaluation(studentID string, criterion *Criterion, scoreValue float64) (*Evaluation, error) {
	if math.IsNaN(scoreValue) || math.IsInf(scoreValue, 0) {
		return nil, fmt.Errorf("оценка содержит недопустимое числовое значение (NaN или Inf)")
	}
	if scoreValue < 0 || scoreValue > criterion.MaxScore {
		return nil, fmt.Errorf("оценка %.2f выходит за пределы криетрия '%s' (0 - %.2f)", scoreValue, criterion.Name, criterion.MaxScore)
	}
	return &Evaluation{
		studentID:   studentID,
		criterionID: criterion.ID,
		score:       Score(scoreValue),
	}, nil
}

func (e *Evaluation) StudentID() string {
	return e.studentID
}

func (e *Evaluation) CriterionID() string {
	return e.criterionID
}

func (e *Evaluation) ScoreValue() float64 {
	return float64(e.score)
}

func (e *Evaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"StudentID":   e.studentID,
		"CriterionID": e.criterionID,
		"Score":       e.score,
	})
}

func (e *Evaluation) UnmarshalJSON(data []byte) error {
	var aux struct {
		StudentID   string  `json:"StudentID"`
		CriterionID string  `json:"CriterionID"`
		Score       float64 `json:"Score"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	e.studentID = aux.StudentID
	e.criterionID = aux.CriterionID
	e.score = Score(aux.Score)
	return nil
}
