package recipeRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/recipeEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicineRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/patientRepo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	ID              uint                  `gorm:"primary_key:auto_increment"`
	Uuid            uuid.UUID             `gorm:"type:varchar(255)"`
	MedicalRecordID uint                  `gorm:"type:uint"`
	MedicineID      uint                  `gorm:"type:uint"`
	Medicine        medicineRepo.Medicine `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	PatientID       uint                  `gorm:"type:uint"`
	Patient         patientRepo.Patient   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	UserID          uint                  `gorm:"type:uint"`
	Username        string                `gorm:"type:varchar"`
	UserRole        string                `gorm:"type:varchar"`
	ConsumptionRule string                `gorm:"type:varchar"`
	Symptom         string                `gorm:"type:varchar"`
	Record          string                `gorm:"type:varchar"`
}

func ToDomain(rec *Recipe) recipeEntity.Domain {
	return recipeEntity.Domain{
		ID:              rec.ID,
		Uuid:            rec.Uuid,
		MedicalRecordID: rec.MedicalRecordID,
		MedicineID:      rec.MedicineID,
		Medicine:        rec.Medicine.Name,
		PatientID:       rec.PatientID,
		Patient:         rec.Patient.Name,
		UserID:          rec.UserID,
		Username:        rec.Username,
		UserRole:        rec.UserRole,
		ConsumptionRule: rec.ConsumptionRule,
		Symptom:         rec.Symptom,
		Record:          rec.Record,
	}
}

func FromDomain(recipeDomain *recipeEntity.Domain) *Recipe {
	return &Recipe{
		ID:              recipeDomain.ID,
		Uuid:            recipeDomain.Uuid,
		MedicalRecordID: recipeDomain.MedicalRecordID,
		MedicineID:      recipeDomain.MedicineID,
		PatientID:       recipeDomain.PatientID,
		UserID:          recipeDomain.UserID,
		Username:        recipeDomain.Username,
		UserRole:        recipeDomain.UserRole,
		ConsumptionRule: recipeDomain.ConsumptionRule,
		Symptom:         recipeDomain.Symptom,
		Record:          recipeDomain.Record,
	}
}

func toDomainArray(record []Recipe) []recipeEntity.Domain {
	var res []recipeEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
