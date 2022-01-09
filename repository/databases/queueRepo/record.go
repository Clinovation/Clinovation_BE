package queueRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/queueEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Queue struct {
	gorm.Model
	ID         uint      `gorm:"primary_key:auto_increment"`
	Uuid       uuid.UUID `gorm:"type:varchar(255)"`
	UserID     uint      `gorm:"type:uint"`
	PatientID  uint      `gorm:"type:uint"`
	ScheduleID uint      `gorm:"type:uint"`
	Role       string    `gorm:"type:varchar(20)"`
}

func ToDomain(rec *Queue) queueEntity.Domain {
	return queueEntity.Domain{
		ID:         rec.ID,
		Uuid:       rec.Uuid,
		PatientID:  rec.PatientID,
		ScheduleID: rec.ScheduleID,
		UserID:     rec.UserID,
		Role:       rec.Role,
	}
}

func FromDomain(queueDomain *queueEntity.Domain) *Queue {
	return &Queue{
		ID:         queueDomain.ID,
		Uuid:       queueDomain.Uuid,
		PatientID:  queueDomain.PatientID,
		UserID:     queueDomain.UserID,
		ScheduleID: queueDomain.ScheduleID,
		Role:       queueDomain.Role,
	}
}

func toDomainArray(record []Queue) []queueEntity.Domain {
	var res []queueEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
