package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/google/uuid"
	"time"
)

type MedicalStaff struct {
	Uuid           uuid.UUID `json:"uuid"`
	Nik            string    `json:"nik"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Dob            string    `json:"dob"`
	Sex            string    `json:"sex"`
	Contact        string    `json:"contact"`
	WorkExperience string    `json:"work_experience"`
	Avatar         string    `json:"avatar"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FromDomain(domain *medicalStaffEntity.Domain) *MedicalStaff {
	return &MedicalStaff{
		Uuid:           domain.Uuid,
		Nik:            domain.Nik,
		Name:           domain.Name,
		Email:          domain.Email,
		Dob:            domain.Dob,
		Sex:            domain.Sex,
		Contact:        domain.Contact,
		WorkExperience: domain.WorkExperience,
		Avatar:         domain.Avatar,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func FromDomainArray(domain []medicalStaffEntity.Domain) []MedicalStaff {
	var res []MedicalStaff
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}

type MedicalStaffLogin struct {
	Token string `json:"token"`
}
