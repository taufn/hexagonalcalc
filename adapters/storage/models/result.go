package models

import (
	"hexagonalcalc/domain"
	"time"
)

type Result struct {
	ID        string
	Left      int
	Right     int
	Result    int
	CreatedAt time.Time
}

func ResultFromDomainModel(params domain.Result) Result {
	return Result{
		Left:   params.Left,
		Right:  params.Right,
		Result: params.Result,
	}
}

func (result Result) ToDomainModel() domain.Result {
	return domain.Result{
		Left:   result.Left,
		Right:  result.Right,
		Result: result.Result,
	}
}
