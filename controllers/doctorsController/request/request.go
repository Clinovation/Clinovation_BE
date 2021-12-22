package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"

	"github.com/google/uuid"
)

type DoctorRegistration struct {
	Name        string `json:"name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password" validate:"required,password"`
}

type DoctorLogin struct {
	Uuid     uuid.UUID `json:"uuid"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,password"`
}

type DoctorUploadAvatar struct {
	Avatar string `json:"avatar"`
}

func (rec *DoctorRegistration) ToDomain() *doctorsEntity.Domain {
	return &doctorsEntity.Domain{
		Name:        rec.Name,
		Username:    rec.Username,
		Email:       rec.Email,
		PhoneNumber: rec.PhoneNumber,
		Avatar:      rec.Avatar,
		Password:    rec.Password,
	}
}
