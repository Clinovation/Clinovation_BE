package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"time"

	"github.com/google/uuid"
)

type Doctors struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created_at"`
}

func FromDomain(domain *doctorsEntity.Domain) *Doctors {
	return &Doctors{
		Uuid:        domain.Uuid,
		Name:        domain.Name,
		Email:       domain.Email,
		PhoneNumber: domain.PhoneNumber,
		Avatar:      domain.Avatar,
		CreatedAt:   domain.CreatedAt,
	}
}

type DoctorsLogin struct {
	Token string `json:"token"`
}
