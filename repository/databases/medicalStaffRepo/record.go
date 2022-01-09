package medicalStaffRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalStaff struct {
	gorm.Model
	ID             uint      `gorm:"primary_key:auto_increment"`
	Uuid           uuid.UUID `gorm:"type:varchar(255)"`
	Nik            string    `gorm:"type:varchar(16)"`
	Name           string    `gorm:"type:varchar(255)"`
	Email          string    `gorm:"uniqueIndex;type:varchar(255)"`
	Dob            string    `gorm:"type:varchar(50)"`
	Sex            string    `gorm:"type:varchar(6)"`
	Contact        string    `gorm:"type:varchar(15)"`
	Password       string    `gorm:"->;<-;not null" `
	WorkExperience string    `gorm:"type:varchar(255)"`
	Avatar         string    `gorm:"type:varchar(255)"`
	Role           string    `gorm:"type:varchar(20)"`
}

func ToDomain(rec *MedicalStaff) medicalStaffEntity.Domain {
	return medicalStaffEntity.Domain{
		ID:             rec.ID,
		Uuid:           rec.Uuid,
		Nik:            rec.Nik,
		Name:           rec.Name,
		Email:          rec.Email,
		Dob:            rec.Dob,
		Sex:            rec.Sex,
		Contact:        rec.Contact,
		Password:       rec.Password,
		WorkExperience: rec.WorkExperience,
		Avatar:         rec.Avatar,
		Role:           rec.Role,
	}
}

func FromDomain(medicalStaffDomain *medicalStaffEntity.Domain) *MedicalStaff {
	return &MedicalStaff{
		ID:             medicalStaffDomain.ID,
		Uuid:           medicalStaffDomain.Uuid,
		Nik:            medicalStaffDomain.Nik,
		Name:           medicalStaffDomain.Name,
		Email:          medicalStaffDomain.Email,
		Dob:            medicalStaffDomain.Dob,
		Sex:            medicalStaffDomain.Sex,
		Contact:        medicalStaffDomain.Contact,
		Password:       medicalStaffDomain.Password,
		WorkExperience: medicalStaffDomain.WorkExperience,
		Avatar:         medicalStaffDomain.Avatar,
		Role:           medicalStaffDomain.Role,
	}
}

func toDomainArray(record []MedicalStaff) []medicalStaffEntity.Domain {
	var res []medicalStaffEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
