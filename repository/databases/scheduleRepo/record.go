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
	UserRole   string    `gorm:"type:varchar(20)"`
}

func ToDomain(rec *Schedule) scheduleEntity.Domain {
	return scheduleEntity.Domain{
		ID:         rec.ID,
		Uuid:       rec.Uuid,
		WorkDayID:  rec.WorkDayID,
		WorkHourID: rec.WorkHourID,
		UserID:     rec.UserID,
		UserRole:   rec.UserRole,
	}
}

func FromDomain(scheduleDomain *scheduleEntity.Domain) *Schedule {
	return &Schedule{
		ID:         scheduleDomain.ID,
		Uuid:       scheduleDomain.Uuid,
		WorkDayID:  scheduleDomain.WorkDayID,
		UserID:     scheduleDomain.UserID,
		WorkHourID: scheduleDomain.WorkHourID,
		UserRole:   scheduleDomain.UserRole,
	}
}

func toDomainArray(record []Schedule) []scheduleEntity.Domain {
	var res []scheduleEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
