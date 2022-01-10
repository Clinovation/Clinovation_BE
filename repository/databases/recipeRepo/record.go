package recipeRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/recipeEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	ID              uint      `gorm:"primary_key:auto_increment"`
	Uuid            uuid.UUID `gorm:"type:varchar(255)"`
	MedicineID      uint      `gorm:"type:uint"`
	ConsumptionRule string    `gorm:"type:varchar(255)"`
}

func ToDomain(rec *Recipe) recipeEntity.Domain {
	return recipeEntity.Domain{
		ID:              rec.ID,
		Uuid:            rec.Uuid,
		MedicineID:      rec.MedicineID,
		ConsumptionRule: rec.ConsumptionRule,
	}
}

func FromDomain(recipeDomain *recipeEntity.Domain) *Recipe {
	return &Recipe{
		ID:              recipeDomain.ID,
		Uuid:            recipeDomain.Uuid,
		MedicineID:      recipeDomain.MedicineID,
		ConsumptionRule: recipeDomain.ConsumptionRule,
	}
}

func toDomainArray(record []Recipe) []recipeEntity.Domain {
	var res []recipeEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
