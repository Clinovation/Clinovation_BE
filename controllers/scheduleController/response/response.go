package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/scheduleEntity"
	"github.com/google/uuid"
	"time"
)

type Schedules struct {
	Uuid       uuid.UUID `json:"uuid"`
	UserID     uint      `json:"user_id"`
	WorkDayID  uint      `json:"work_day_id"`
	WorkHourID uint      `json:"work_hour_id"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain *scheduleEntity.Domain) *Schedules {
	return &Schedules{
		Uuid:       domain.Uuid,
		UserID:     domain.UserID,
		WorkDayID:  domain.WorkDayID,
		WorkHourID: domain.WorkHourID,
		Role:       domain.Role,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomainArray(domain []scheduleEntity.Domain) []Schedules {
	var res []Schedules
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}
