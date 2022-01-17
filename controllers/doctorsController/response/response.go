package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"time"

	"github.com/google/uuid"
)

type Doctors struct {
	Uuid           uuid.UUID `json:"uuid"`
	Nik            string    `json:"nik"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Dob            string    `json:"dob"`
	Sex            string    `json:"sex"`
	Contact        string    `json:"contact"`
	Specialist     string    `json:"specialist"`
	WorkExperience string    `json:"work_experience"`
	Avatar         string    `json:"avatar"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FromDomain(domain *doctorsEntity.Domain) *Doctors {
	return &Doctors{
		Uuid:           domain.Uuid,
		Nik:            domain.Nik,
		Name:           domain.Name,
		Email:          domain.Email,
		Dob:            domain.Dob,
		Sex:            domain.Sex,
		Contact:        domain.Contact,
		Specialist:     domain.Specialist,
		WorkExperience: domain.WorkExperience,
		Avatar:         domain.Avatar,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func FromDomainArray(domain []doctorsEntity.Domain) []Doctors {
	var res []Doctors
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}

type DoctorsLogin struct {
	Token string `json:"token"`
}
