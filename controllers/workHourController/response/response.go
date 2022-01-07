package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/workHourEntity"
	"time"

	"github.com/google/uuid"
)

type WorkHours struct {
	Uuid      uuid.UUID `json:"uuid"`
	Hour      string    `json:"hour"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain *workHourEntity.Domain) *WorkHours {
	return &WorkHours{
		Uuid:      domain.Uuid,
		Hour:      domain.Hour,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainArray(domain []workHourEntity.Domain) []WorkHours {
	var res []WorkHours
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}
