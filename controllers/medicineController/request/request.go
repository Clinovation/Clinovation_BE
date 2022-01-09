package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicineEntity"
)

type MedicineRegistration struct {
	Name  string `json:"name" validate:"required"`
	Type  string `json:"type" validate:"required"`
	Price int    `json:"price" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}

func (rec *MedicineRegistration) ToDomain() *medicineEntity.Domain {
	return &medicineEntity.Domain{
		Type:  rec.Type,
		Price: rec.Price,
		Stock: rec.Stock,
		Name:  rec.Name,
	}
}
