package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
)

type PatientRegistration struct {
	Nik            string `json:"nik" validate:"required"`
	Name           string `json:"name" validate:"required"`
	MedicalStaffID uint   `json:"medical_staff_id"`
	Dob            string `json:"dob" validate:"required"`
	Sex            string `json:"sex" validate:"required"`
	Contact        string `json:"contact" validate:"required"`
	StatusMartial  string `json:"status_martial" validate:"required"`
	Address        string `json:"address" validate:"required"`
	Height         string `json:"height" validate:"required"`
	Weight         string `json:"weight" validate:"required"`
	Record         string `json:"record" validate:"required"`
	Avatar         string `json:"avatar"`
}

type PatientUpdate struct {
	Nik            string `json:"nik" `
	Name           string `json:"name" `
	MedicalStaffID uint   `json:"medical_staff_id"`
	Dob            string `json:"dob" `
	Sex            string `json:"sex" `
	Contact        string `json:"contact" `
	StatusMartial  string `json:"status_martial" `
	Address        string `json:"address" `
	Height         string `json:"height" `
	Weight         string `json:"weight" `
	Record         string `json:"record" `
	Avatar         string `json:"avatar"`
}

type PatientUploadAvatar struct {
	Avatar string `json:"avatar" validate:"required"`
}

func (rec *PatientRegistration) ToDomain() *patientEntity.Domain {
	return &patientEntity.Domain{
		Nik:            rec.Nik,
		MedicalStaffID: rec.MedicalStaffID,
		Name:           rec.Name,
		Dob:            rec.Dob,
		Sex:            rec.Sex,
		Contact:        rec.Contact,
		StatusMartial:  rec.StatusMartial,
		Address:        rec.Address,
		Height:         rec.Height,
		Weight:         rec.Weight,
		Record:         rec.Record,
		Avatar:         rec.Avatar,
	}
}

func (rec *PatientUpdate) ToDomainUpdate() *patientEntity.Domain {
	return &patientEntity.Domain{
		Nik:            rec.Nik,
		MedicalStaffID: rec.MedicalStaffID,
		Name:           rec.Name,
		Dob:            rec.Dob,
		Sex:            rec.Sex,
		Contact:        rec.Contact,
		StatusMartial:  rec.StatusMartial,
		Address:        rec.Address,
		Height:         rec.Height,
		Weight:         rec.Weight,
		Record:         rec.Record,
		Avatar:         rec.Avatar,
	}
}
