package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
)

type MedicalRecordRegistration struct {
	Consultation string `json:"consultation" validate:"required"`
	NewRecord    string `json:"new_record" `
}

func (rec *MedicalRecordRegistration) ToDomain() *medicalRecordEntity.Domain {
	return &medicalRecordEntity.Domain{
		Consultation: rec.Consultation,
		NewRecord:    rec.NewRecord,
	}
}
