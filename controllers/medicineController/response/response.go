package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicineEntity"
	"github.com/google/uuid"
	"time"
)

type Medicine struct {
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain *medicineEntity.Domain) *Medicine {
	return &Medicine{
		Uuid:      domain.Uuid,
		Name:      domain.Name,
		Type:      domain.Type,
		Price:     domain.Price,
		Stock:     domain.Stock,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainArray(domain []medicineEntity.Domain) []Medicine {
	var res []Medicine
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}
