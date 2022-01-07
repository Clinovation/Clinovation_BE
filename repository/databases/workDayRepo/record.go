package workDayRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkDays struct {
	gorm.Model
	ID   uint      `gorm:"primary_key:auto_increment"`
	Uuid uuid.UUID `gorm:"type:varchar(255)"`
	Day  string    `gorm:"type:varchar(30)"`
}

func ToDomain(rec *WorkDays) workDayEntity.Domain {
	return workDayEntity.Domain{
		ID:   rec.ID,
		Uuid: rec.Uuid,
		Day:  rec.Day,
	}
}

func FromDomain(workDayDomain *workDayEntity.Domain) *WorkDays {
	return &WorkDays{
		ID:   workDayDomain.ID,
		Uuid: workDayDomain.Uuid,
		Day:  workDayDomain.Day,
	}
}

func toDomainArray(record []WorkDays) []workDayEntity.Domain {
	var res []workDayEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
