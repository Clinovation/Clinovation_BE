package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/scheduleEntity"
)

type ScheduleRegistration struct {
	UserID     uint   `json:"user_id"`
	WorkDayID  uint   `json:"work_day_id"`
	WorkHourID uint   `json:"work_hour_id"`
	Role       string `json:"role"`
}

func (rec *ScheduleRegistration) ToDomain() *scheduleEntity.Domain {
	return &scheduleEntity.Domain{
		UserID:     rec.UserID,
		WorkDayID:  rec.WorkDayID,
		WorkHourID: rec.WorkHourID,
		Role:       rec.Role,
	}
}
