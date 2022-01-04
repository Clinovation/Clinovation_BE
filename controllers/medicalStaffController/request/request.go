package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/google/uuid"
)

type MedicalStaffRegistration struct {
	Nik            string `json:"nik" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Dob            string `json:"dob" validate:"required"`
	Sex            string `json:"sex" validate:"required"`
	Contact        string `json:"contact" validate:"required"`
	Password       string `json:"password" validate:"required,password"`
	WorkExperience string `json:"work_experience" validate:"required"`
	Avatar         string `json:"avatar"`
}

type MedicalStaffLogin struct {
	Uuid     uuid.UUID `json:"uuid"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,password"`
}

type MedicalStaffUploadAvatar struct {
	Avatar string `json:"avatar" validate:"required"`
}

func (rec *MedicalStaffRegistration) ToDomain() *medicalStaffEntity.Domain {
	return &medicalStaffEntity.Domain{
		Nik:            rec.Nik,
		Name:           rec.Name,
		Email:          rec.Email,
		Dob:            rec.Dob,
		Sex:            rec.Sex,
		Contact:        rec.Contact,
		Password:       rec.Password,
		WorkExperience: rec.WorkExperience,
		Avatar:         rec.Avatar,
	}
}
