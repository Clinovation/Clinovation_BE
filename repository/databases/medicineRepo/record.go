package medicineRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicineEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Medicine struct {
	gorm.Model
	ID    uint      `gorm:"primary_key:auto_increment"`
	Uuid  uuid.UUID `gorm:"type:varchar(255)"`
	Name  string    `gorm:"type:varchar(255)"`
	Type  string    `gorm:"type:varchar(30)"`
	Price int       `gorm:"type:integer"`
	Stock int       `gorm:"type:integer"`
}

func ToDomain(rec *Medicine) medicineEntity.Domain {
	return medicineEntity.Domain{
		ID:    rec.ID,
		Uuid:  rec.Uuid,
		Name:  rec.Name,
		Type:  rec.Type,
		Price: rec.Price,
		Stock: rec.Stock,
	}
}

func FromDomain(medicineDomain *medicineEntity.Domain) *Medicine {
	return &Medicine{
		ID:    medicineDomain.ID,
		Uuid:  medicineDomain.Uuid,
		Name:  medicineDomain.Name,
		Type:  medicineDomain.Type,
		Price: medicineDomain.Price,
		Stock: medicineDomain.Stock,
	}
}

func toDomainArray(record []Medicine) []medicineEntity.Domain {
	var res []medicineEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
