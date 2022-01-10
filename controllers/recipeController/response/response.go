package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/recipeEntity"
	"github.com/google/uuid"
	"time"
)

type Recipe struct {
	Uuid            uuid.UUID `json:"uuid"`
	MedicineID      uint      `json:"medicine_id"`
	ConsumptionRule string    `json:"consumption_rule"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func FromDomain(domain *recipeEntity.Domain) *Recipe {
	return &Recipe{
		Uuid:            domain.Uuid,
		MedicineID:      domain.MedicineID,
		ConsumptionRule: domain.ConsumptionRule,
	}
}

func FromDomainArray(domain []recipeEntity.Domain) []Recipe {
	var res []Recipe
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}
