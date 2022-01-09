package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/queueEntity"
	"github.com/google/uuid"
	"time"
)

type Queues struct {
	Uuid       uuid.UUID `json:"uuid"`
	UserID     uint      `json:"user_id"`
	ScheduleID uint      `json:"schedule_id"`
	PatientID  uint      `json:"patient_id"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain *queueEntity.Domain) *Queues {
	return &Queues{
		Uuid:       domain.Uuid,
		UserID:     domain.UserID,
		ScheduleID: domain.ScheduleID,
		PatientID:  domain.PatientID,
		Role:       domain.Role,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomainArray(domain []queueEntity.Domain) []Queues {
	var res []Queues
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}
