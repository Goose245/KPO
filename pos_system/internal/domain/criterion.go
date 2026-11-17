package domain

import (
	"errors"
	"math"
)

type Criterion struct {
	ID       string
	Name     string
	MaxScore float64
	Weight   float64
}

func NewCriterion(id, name string, maxScore, weight float64) (*Criterion, error) {
	if math.IsNaN(maxScore) || math.IsInf(maxScore, 0) || math.IsNaN(weight) || math.IsInf(weight, 0) {
		return nil, errors.New("недопустимые числовые значения (NaN или Inf)")
	}
	if maxScore <= 0 {
		return nil, errors.New("макс бал должен быть больше 0")
	}
	if weight <= 0 {
		return nil, errors.New("вес критерия должен быть положительным")
	}
	return &Criterion{
		ID:       id,
		Name:     name,
		MaxScore: maxScore,
		Weight:   weight,
	}, nil
}
