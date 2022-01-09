package scheduleRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/scheduleEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	ID         uint      `gorm:"primary_key:auto_increment"`
	Uuid       uuid.UUID `gorm:"type:varchar(255)"`
	UserID     uint      `gorm:"type:uint"`
	WorkDayID  uint      `gorm:"type:uint"`
	WorkHourID uint      `gorm:"type:uint"`
	Role       string    `gorm:"type:varchar(20)"`
}

func ToDomain(rec *Schedule) scheduleEntity.Domain {
	return scheduleEntity.Domain{
		ID:         rec.ID,
		Uuid:       rec.Uuid,
		WorkDayID:  rec.WorkDayID,
		WorkHourID: rec.WorkHourID,
		UserID:     rec.UserID,
		Role:       rec.Role,
	}
}

func FromDomain(doctorDomain *scheduleEntity.Domain) *Schedule {
	return &Schedule{
		ID:         doctorDomain.ID,
		Uuid:       doctorDomain.Uuid,
		WorkDayID:  doctorDomain.WorkDayID,
		UserID:     doctorDomain.UserID,
		WorkHourID: doctorDomain.WorkHourID,
		Role:       doctorDomain.Role,
	}
}

func toDomainArray(record []Schedule) []scheduleEntity.Domain {
	var res []scheduleEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
