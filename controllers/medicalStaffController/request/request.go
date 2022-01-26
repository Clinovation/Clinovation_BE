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

type MedicalStaffUpdate struct {
	Nik            string `json:"nik" `
	Name           string `json:"name" `
	Email          string `json:"email" `
	Dob            string `json:"dob" `
	Sex            string `json:"sex" `
	Contact        string `json:"contact" `
	Password       string `json:"password" `
	WorkExperience string `json:"work_experience" `
	Avatar         string `json:"avatar"`
}

type MedicalStaffLogin struct {
	Uuid     uuid.UUID `json:"uuid"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,password"`
}

type ForgetPassword struct {
	Email string `json:"email" validate:"required,email"`
	Nik   string `json:"nik" validate:"required"`
}

type ChangePassword struct {
	Password string `json:"password" validate:"required,password"`
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

func (rec *MedicalStaffUpdate) ToDomainUpdate() *medicalStaffEntity.Domain {
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

func (rec *ForgetPassword) ToDomainForget() *medicalStaffEntity.Domain {
	return &medicalStaffEntity.Domain{
		Nik:   rec.Nik,
		Email: rec.Email,
	}
}

func (rec *ChangePassword) ToDomainChange() *medicalStaffEntity.Domain {
	return &medicalStaffEntity.Domain{
		Password: rec.Password,
	}
}
