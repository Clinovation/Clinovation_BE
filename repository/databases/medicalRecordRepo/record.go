package medicalRecordRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecord struct {
	gorm.Model
	ID           uint      `gorm:"primary_key:auto_increment"`
	Uuid         uuid.UUID `gorm:"type:varchar(255)"`
	UserID       uint      `gorm:"type:uint"`
	PatientID    uint      `gorm:"type:uint"`
	RecipeID     uint      `gorm:"type:uint"`
	Consultation string    `gorm:"type:uint"`
	Symptom      string    `gorm:"type:uint"`
	Note         string    `gorm:"type:uint"`
}

func ToDomain(rec *MedicalRecord) medicalRecordEntity.Domain {
	return medicalRecordEntity.Domain{
		ID:           rec.ID,
		Uuid:         rec.Uuid,
		PatientID:    rec.PatientID,
		RecipeID:     rec.RecipeID,
		UserID:       rec.UserID,
		Consultation: rec.Consultation,
		Symptom:      rec.Symptom,
		Note:         rec.Note,
	}
}

func FromDomain(medicalRecordDomain *medicalRecordEntity.Domain) *MedicalRecord {
	return &MedicalRecord{
		ID:           medicalRecordDomain.ID,
		Uuid:         medicalRecordDomain.Uuid,
		PatientID:    medicalRecordDomain.PatientID,
		UserID:       medicalRecordDomain.UserID,
		RecipeID:     medicalRecordDomain.RecipeID,
		Consultation: medicalRecordDomain.Consultation,
		Symptom:      medicalRecordDomain.Symptom,
		Note:         medicalRecordDomain.Note,
	}
}

func toDomainArray(record []MedicalRecord) []medicalRecordEntity.Domain {
	var res []medicalRecordEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
