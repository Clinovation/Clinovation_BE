package workHourRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/workHourEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkHours struct {
	gorm.Model
	ID   uint      `gorm:"primary_key:auto_increment"`
	Uuid uuid.UUID `gorm:"type:varchar(255)"`
	Hour string    `gorm:"type:varchar(30)"`
}

func ToDomain(rec *WorkHours) workHourEntity.Domain {
	return workHourEntity.Domain{
		ID:   rec.ID,
		Uuid: rec.Uuid,
		Hour: rec.Hour,
	}
}

func FromDomain(workHourDomain *workHourEntity.Domain) *WorkHours {
	return &WorkHours{
		ID:   workHourDomain.ID,
		Uuid: workHourDomain.Uuid,
		Hour: workHourDomain.Hour,
	}
}

func toDomainArray(record []WorkHours) []workHourEntity.Domain {
	var res []workHourEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
