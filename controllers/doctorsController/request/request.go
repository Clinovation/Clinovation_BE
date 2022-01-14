package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"

	"github.com/google/uuid"
)

type DoctorRegistration struct {
	Nik            string `json:"nik" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Dob            string `json:"dob" validate:"required"`
	Sex            string `json:"sex" validate:"required"`
	Contact        string `json:"contact" validate:"required"`
	Password       string `json:"password" validate:"required,password"`
	Specialist     string `json:"specialist" validate:"required"`
	WorkExperience string `json:"work_experience" validate:"required"`
	Avatar         string `json:"avatar"`
}

type DoctorLogin struct {
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

type DoctorUploadAvatar struct {
	Avatar string `json:"avatar" validate:"required"`
}

func (rec *DoctorRegistration) ToDomain() *doctorsEntity.Domain {
	return &doctorsEntity.Domain{
		Nik:            rec.Nik,
		Name:           rec.Name,
		Email:          rec.Email,
		Dob:            rec.Dob,
		Sex:            rec.Sex,
		Contact:        rec.Contact,
		Password:       rec.Password,
		Specialist:     rec.Specialist,
		WorkExperience: rec.WorkExperience,
		Avatar:         rec.Avatar,
	}
}

func (rec *ForgetPassword) ToDomainForget() *doctorsEntity.Domain {
	return &doctorsEntity.Domain{
		Nik:   rec.Nik,
		Email: rec.Email,
	}
}

func (rec *ChangePassword) ToDomainChange() *doctorsEntity.Domain {
	return &doctorsEntity.Domain{
		Password: rec.Password,
	}
}
