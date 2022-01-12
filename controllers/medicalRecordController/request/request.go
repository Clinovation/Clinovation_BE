package request

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
)

type MedicalRecordRegistration struct {
	UserID       uint   `json:"user_id"`
	PatientID    uint   `json:"patient_id"`
	RecipeID     uint   `json:"recipe_id"`
	Consultation string `json:"consultation"`
	Symptom      string `json:"symptom"`
	NewNote      string `json:"new_note"`
}

func (rec *MedicalRecordRegistration) ToDomain() *medicalRecordEntity.Domain {
	return &medicalRecordEntity.Domain{
		UserID:       rec.UserID,
		PatientID:    rec.PatientID,
		RecipeID:     rec.RecipeID,
		Consultation: rec.Consultation,
		Symptom:      rec.Symptom,
		NewNote:      rec.NewNote,
	}
}
