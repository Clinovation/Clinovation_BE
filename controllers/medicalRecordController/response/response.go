package response

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/google/uuid"
	"time"
)

type MedicalRecords struct {
	Uuid                 uuid.UUID `json:"uuid"`
	PatientID            uint      `json:"patient_id"`
	PatientName          string    `json:"patient_name"`
	PatientAddress       string    `json:"patient_address"`
	PatientDob           string    `json:"patient_dob"`
	PatientHeight        string    `json:"patient_height"`
	PatientNik           string    `json:"patient_nik"`
	PatientSex           string    `json:"patient_sex"`
	PatientStatusMartial string    `json:"patient_status_martial"`
	PatientWeight        string    `json:"patient_weight"`
	PatientUuid          string    `json:"patient_uuid"`
	PatientAvatar        string    `json:"patient_avatar"`
	UserID               uint      `json:"user_id"`
	Username             string    `json:"username"`
	UserRole             string    `json:"user_role"`
	UserSpecialist       string    `json:"user_specialist"`
	MedicalStaffID       uint      `json:"medical_staff_id"`
	MedicalStaff         string    `json:"medical_staff"`
	Consultation         string    `json:"consultation"`
	NewRecord            string    `json:"new_record"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func FromDomain(domain *medicalRecordEntity.Domain) *MedicalRecords {
	return &MedicalRecords{
		Uuid:                 domain.Uuid,
		PatientID:            domain.PatientID,
		PatientName:          domain.PatientName,
		PatientAddress:       domain.PatientAddress,
		PatientDob:           domain.PatientDob,
		PatientHeight:        domain.PatientHeight,
		PatientNik:           domain.PatientNik,
		PatientSex:           domain.PatientSex,
		PatientStatusMartial: domain.PatientStatusMartial,
		PatientWeight:        domain.PatientWeight,
		PatientAvatar:        domain.PatientAvatar,
		PatientUuid:          domain.PatientUuid,
		UserID:               domain.UserID,
		Username:             domain.Username,
		UserRole:             domain.UserRole,
		UserSpecialist:       domain.UserSpecialist,
		MedicalStaffID:       domain.MedicalStaffID,
		MedicalStaff:         domain.MedicalStaff,
		Consultation:         domain.Consultation,
		NewRecord:            domain.NewRecord,
		CreatedAt:            domain.CreatedAt,
		UpdatedAt:            domain.UpdatedAt,
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
