package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/recipeEntity"
)

type RecipeRegistration struct {
	ConsumptionRule string `json:"consumption_rule" validate:"required"`
}

func (rec *RecipeRegistration) ToDomain() *recipeEntity.Domain {
	return &recipeEntity.Domain{
		ConsumptionRule: rec.ConsumptionRule,
	}
}
