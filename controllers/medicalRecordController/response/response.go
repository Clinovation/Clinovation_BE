package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/google/uuid"
	"time"
)

type MedicalRecords struct {
	Uuid         uuid.UUID `json:"uuid"`
	UserID       uint      `json:"user_id"`
	RecipeID     uint      `json:"recipe_id"`
	PatientID    uint      `json:"patient_id"`
	Consultation string    `json:"consultation"`
	Symptom      string    `json:"symptom"`
	Note         string    `json:"note"`
	NewNote      string    `json:"new_note"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(domain *medicalRecordEntity.Domain) *MedicalRecords {
	return &MedicalRecords{
		Uuid:         domain.Uuid,
		UserID:       domain.UserID,
		RecipeID:     domain.RecipeID,
		PatientID:    domain.PatientID,
		Consultation: domain.Consultation,
		Symptom:      domain.Symptom,
		Note:         domain.Note,
		NewNote:      domain.NewNote,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func FromDomainArray(domain []medicalRecordEntity.Domain) []MedicalRecords {
	var res []MedicalRecords
	for _, v := range domain {
		res = append(res, *FromDomain(&v))
	}
	return res
}
