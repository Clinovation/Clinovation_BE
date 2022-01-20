package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
	"time"

	"github.com/google/uuid"
)

type WorkDays struct {
	Uuid      uuid.UUID `json:"uuid"`
	Day       string    `json:"day"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain *workDayEntity.Domain) *WorkDays {
	return &WorkDays{
		Uuid:      domain.Uuid,
		Day:       domain.Day,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainArray(domain []workDayEntity.Domain) []WorkDays {
	var res []WorkDays
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}

