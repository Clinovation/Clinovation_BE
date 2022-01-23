package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/google/uuid"
	"time"
)

type MedicalRecords struct {
	Uuid           uuid.UUID `json:"uuid"`
	PatientID      uint      `json:"patient_id"`
	Patient        string    `json:"patient"`
	UserID         uint      `json:"user_id"`
	Username       string    `json:"username"`
	UserRole       string    `json:"user_role"`
	UserSpecialist string    `json:"user_specialist"`
	MedicalStaffID uint      `json:"medical_staff_id"`
	MedicalStaff   string    `json:"medical_staff"`
	Consultation   string    `json:"consultation"`
	NewRecord      string    `json:"new_record"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FromDomain(domain *medicalRecordEntity.Domain) *MedicalRecords {
	return &MedicalRecords{
		Uuid:           domain.Uuid,
		PatientID:      domain.PatientID,
		Patient:        domain.Patient,
		UserID:         domain.UserID,
		Username:       domain.Username,
		UserRole:       domain.UserRole,
		UserSpecialist: domain.UserSpecialist,
		MedicalStaffID: domain.MedicalStaffID,
		MedicalStaff:   domain.MedicalStaff,
		Consultation:   domain.Consultation,
		NewRecord:      domain.NewRecord,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func FromDomainArray(domain []medicalRecordEntity.Domain) []MedicalRecords {
	var res []MedicalRecords
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
