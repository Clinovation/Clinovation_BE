package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
)

type WorkDayRegistration struct {
	Day string `json:"day" validate:"required"`
}

func (rec *WorkDayRegistration) ToDomain() *workDayEntity.Domain {
	return &workDayEntity.Domain{
		Day: rec.Day,
	}
}
