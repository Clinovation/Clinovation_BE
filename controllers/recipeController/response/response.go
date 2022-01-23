package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/recipeEntity"
	"github.com/google/uuid"
	"time"
)

type Recipe struct {
	Uuid            uuid.UUID `json:"uuid"`
	MedicalRecordID uint      `json:"medical_record_id"`
	MedicineID      uint      `json:"medicine_id"`
	Medicine        string    `json:"medicine"`
	PatientID       uint      `json:"patient_id"`
	Patient         string    `json:"patient"`
	UserID          uint      `json:"user_id"`
	Username        string    `json:"username"`
	UserRole        string    `json:"user_role"`
	ConsumptionRule string    `json:"consumption_rule"`
	Symptom         string    `json:"symptom"`
	Record          string    `json:"record"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func FromDomain(domain *recipeEntity.Domain) *Recipe {
	return &Recipe{
		Uuid:            domain.Uuid,
		MedicalRecordID: domain.MedicalRecordID,
		MedicineID:      domain.MedicineID,
		Medicine:        domain.Medicine,
		PatientID:       domain.PatientID,
		Patient:         domain.Patient,
		UserID:          domain.UserID,
		Username:        domain.Username,
		UserRole:        domain.UserRole,
		ConsumptionRule: domain.ConsumptionRule,
		Symptom:         domain.Symptom,
		Record:          domain.Record,
	}
}

func FromDomainArray(domain []recipeEntity.Domain) []Recipe {
	var res []Recipe
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
