package medicalRecordRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicalStaffRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/patientRepo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecord struct {
	gorm.Model
	ID             uint                          `gorm:"primary_key:auto_increment"`
	Uuid           uuid.UUID                     `gorm:"type:varchar(255)"`
	PatientID      uint                          `gorm:"type:uint"`
	Patient        patientRepo.Patient           `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	UserID         uint                          `gorm:"type:uint"`
	Username       string                        `gorm:"type:varchar"`
	UserRole       string                        `gorm:"type:varchar"`
	UserSpecialist string                        `gorm:"type:varchar"`
	MedicalStaffID uint                          `gorm:"type:uint"`
	MedicalStaff   medicalStaffRepo.MedicalStaff `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Consultation   string                        `gorm:"type:varchar"`
	NewRecord      string                        `gorm:"type:varchar"`
}

func ToDomain(rec *MedicalRecord) medicalRecordEntity.Domain {
	return medicalRecordEntity.Domain{
		ID:             rec.ID,
		Uuid:           rec.Uuid,
		PatientID:      rec.PatientID,
		Patient:        rec.Patient.Name,
		UserID:         rec.UserID,
		Username:       rec.Username,
		UserRole:       rec.UserRole,
		UserSpecialist: rec.UserSpecialist,
		MedicalStaffID: rec.MedicalStaffID,
		MedicalStaff:   rec.MedicalStaff.Name,
		Consultation:   rec.Consultation,
		NewRecord:      rec.NewRecord,
	}
}

func FromDomain(medicalRecordDomain *medicalRecordEntity.Domain) *MedicalRecord {
	return &MedicalRecord{
		ID:             medicalRecordDomain.ID,
		Uuid:           medicalRecordDomain.Uuid,
		PatientID:      medicalRecordDomain.PatientID,
		UserID:         medicalRecordDomain.UserID,
		Username:       medicalRecordDomain.Username,
		UserRole:       medicalRecordDomain.UserRole,
		UserSpecialist: medicalRecordDomain.UserSpecialist,
		MedicalStaffID: medicalRecordDomain.MedicalStaffID,
		Consultation:   medicalRecordDomain.Consultation,
		NewRecord:      medicalRecordDomain.NewRecord,
	}
}

func toDomainArray(record []MedicalRecord) []medicalRecordEntity.Domain {
	var res []medicalRecordEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
