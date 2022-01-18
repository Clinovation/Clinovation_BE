package patientRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicalStaffRepo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID             uint `gorm:"primary_key:auto_increment"`
	MedicalStaffID uint
	MedicalStaff   medicalStaffRepo.MedicalStaff `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Uuid           uuid.UUID                     `gorm:"type:varchar(255)"`
	Nik            string                        `gorm:"type:varchar(16)"`
	Name           string                        `gorm:"type:varchar(255)"`
	Dob            string                        `gorm:"type:varchar(50)"`
	Sex            string                        `gorm:"type:varchar(6)"`
	Contact        string                        `gorm:"type:varchar(15)"`
	StatusMartial  string                        `gorm:"type:varchar(20)"`
	Address        string                        `gorm:"type:varchar(255)"`
	Height         string                        `gorm:"type:varchar(10)"`
	Weight         string                        `gorm:"type:varchar(10)"`
	Record         string                        `gorm:"type:varchar"`
	Avatar         string                        `gorm:"type:varchar(255)"`
	Role           string                        `gorm:"type:varchar(20)"`
}

func ToDomain(rec *Patient) patientEntity.Domain {
	return patientEntity.Domain{
		ID:             rec.ID,
		Uuid:           rec.Uuid,
		MedicalStaffID: rec.MedicalStaffID,
		MedicalStaff:   rec.MedicalStaff.Name,
		Nik:            rec.Nik,
		Name:           rec.Name,
		Dob:            rec.Dob,
		Sex:            rec.Sex,
		Contact:        rec.Contact,
		StatusMartial:  rec.StatusMartial,
		Address:        rec.Address,
		Height:         rec.Height,
		Weight:         rec.Weight,
		Avatar:         rec.Avatar,
		Record:         rec.Record,
		Role:           rec.Role,
	}
}

func FromDomain(patientDomain *patientEntity.Domain) *Patient {
	return &Patient{
		ID:             patientDomain.ID,
		Uuid:           patientDomain.Uuid,
		MedicalStaffID: patientDomain.MedicalStaffID,
		Nik:            patientDomain.Nik,
		Name:           patientDomain.Name,
		Dob:            patientDomain.Dob,
		Sex:            patientDomain.Sex,
		Contact:        patientDomain.Contact,
		StatusMartial:  patientDomain.StatusMartial,
		Address:        patientDomain.Address,
		Height:         patientDomain.Height,
		Weight:         patientDomain.Weight,
		Record:         patientDomain.Record,
		Avatar:         patientDomain.Avatar,
		Role:           patientDomain.Role,
	}
}

func toDomainArray(record []Patient) []patientEntity.Domain {
	var res []patientEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
