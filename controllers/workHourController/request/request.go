package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/workHourEntity"
)

type WorkHourRegistration struct {
	Hour string `json:"hour" validate:"required"`
}

func (rec *WorkHourRegistration) ToDomain() *workHourEntity.Domain {
	return &workHourEntity.Domain{
		Hour: rec.Hour,
	}
}
