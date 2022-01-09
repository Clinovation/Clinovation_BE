package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/queueEntity"
)

type QueueRegistration struct {
	UserID     uint   `json:"user_id"`
	PatientID  uint   `json:"patient_id"`
	ScheduleID uint   `json:"schedule_id"`
	Role       string `json:"role"`
}

func (rec *QueueRegistration) ToDomain() *queueEntity.Domain {
	return &queueEntity.Domain{
		UserID:     rec.UserID,
		PatientID:  rec.PatientID,
		ScheduleID: rec.ScheduleID,
		Role:       rec.Role,
	}
}
