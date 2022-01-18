package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/google/uuid"
	"time"
)

type Patients struct {
	Uuid           uuid.UUID `json:"uuid"`
	MedicalStaffID uint      `json:"medicalStaff_id"`
	MedicalStaff   string    `json:"medical_staff"`
	Nik            string    `json:"nik"`
	Name           string    `json:"name"`
	Email          string    `json:"username"`
	Dob            string    `json:"dob"`
	Sex            string    `json:"sex"`
	Contact        string    `json:"contact"`
	StatusMartial  string    `json:"status_martial"`
	Address        string    `json:"address"`
	Height         string    `json:"height"`
	Weight         string    `json:"weight"`
	Avatar         string    `json:"avatar"`
	Record         string    `json:"record"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FromDomain(domain *patientEntity.Domain) *Patients {
	return &Patients{
		Uuid:           domain.Uuid,
		MedicalStaffID: domain.MedicalStaffID,
		MedicalStaff:   domain.MedicalStaff,
		Nik:            domain.Nik,
		Name:           domain.Name,
		Dob:            domain.Dob,
		Sex:            domain.Sex,
		Contact:        domain.Contact,
		StatusMartial:  domain.StatusMartial,
		Address:        domain.Address,
		Height:         domain.Height,
		Weight:         domain.Weight,
		Avatar:         domain.Avatar,
		Record:         domain.Record,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func FromDomainArray(domain []patientEntity.Domain) []Patients {
	var res []Patients
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
