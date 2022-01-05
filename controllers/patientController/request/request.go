package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
)

type PatientRegistration struct {
	Nik           string `json:"nik" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Dob           string `json:"dob" validate:"required"`
	Sex           string `json:"sex" validate:"required"`
	Contact       string `json:"contact" validate:"required"`
	StatusMartial string `json:"status_martial" validate:"required"`
	Address       string `json:"address" validate:"required"`
	Height        string `json:"height" validate:"required"`
	Weight        string `json:"weight" validate:"required"`
	Avatar        string `json:"avatar"`
}

type PatientUploadAvatar struct {
	Avatar string `json:"avatar" validate:"required"`
}

func (rec *PatientRegistration) ToDomain() *patientEntity.Domain {
	return &patientEntity.Domain{
		Nik:           rec.Nik,
		Name:          rec.Name,
		Dob:           rec.Dob,
		Sex:           rec.Sex,
		Contact:       rec.Contact,
		StatusMartial: rec.StatusMartial,
		Address:       rec.Address,
		Height:        rec.Height,
		Weight:        rec.Weight,
		Avatar:        rec.Avatar,
	}
}
