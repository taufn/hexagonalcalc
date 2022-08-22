package repositories

import (
	"fmt"
	"hexagonalcalc/adapters/storage/models"
	"hexagonalcalc/domain"
	"time"
)

type ResultRepository struct{}

func NewResultRepository() ResultRepository {
	return ResultRepository{}
}

func (r ResultRepository) Save(result domain.Result) (domain.Result, error) {
	model := models.ResultFromDomainModel(result)
	model.ID = fmt.Sprintf("%d", time.Now().Unix())
	model.CreatedAt = time.Now()
	// TODO: perform persistence
	return model.ToDomainModel(), nil
}
