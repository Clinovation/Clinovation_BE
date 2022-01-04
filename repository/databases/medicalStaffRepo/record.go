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

func FromDomain(doctorDomain *medicalStaffEntity.Domain) *MedicalStaff {
	return &MedicalStaff{
		ID:             doctorDomain.ID,
		Uuid:           doctorDomain.Uuid,
		Nik:            doctorDomain.Nik,
		Name:           doctorDomain.Name,
		Email:          doctorDomain.Email,
		Dob:            doctorDomain.Dob,
		Sex:            doctorDomain.Sex,
		Contact:        doctorDomain.Contact,
		Password:       doctorDomain.Password,
		WorkExperience: doctorDomain.WorkExperience,
		Avatar:         doctorDomain.Avatar,
		Role:           doctorDomain.Role,
	}
}

func toDomainArray(record []MedicalStaff) []medicalStaffEntity.Domain {
	var res []medicalStaffEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
