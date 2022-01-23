package medicalRecordRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicalStaffRepo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecord struct {
	gorm.Model
	ID                   uint                          `gorm:"primary_key:auto_increment"`
	Uuid                 uuid.UUID                     `gorm:"type:varchar(255)"`
	PatientID            uint                          `gorm:"type:uint"`
	PatientName          string                        `gorm:"type:varchar"`
	PatientAddress       string                        `gorm:"type:varchar"`
	PatientDob           string                        `gorm:"type:varchar"`
	PatientHeight        string                        `gorm:"type:varchar"`
	PatientNik           string                        `gorm:"type:varchar"`
	PatientSex           string                        `gorm:"type:varchar"`
	PatientStatusMartial string                        `gorm:"type:varchar"`
	PatientWeight        string                        `gorm:"type:varchar"`
	UserID               uint                          `gorm:"type:uint"`
	Username             string                        `gorm:"type:varchar"`
	UserRole             string                        `gorm:"type:varchar"`
	UserSpecialist       string                        `gorm:"type:varchar"`
	MedicalStaffID       uint                          `gorm:"type:uint"`
	MedicalStaff         medicalStaffRepo.MedicalStaff `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Consultation         string                        `gorm:"type:varchar"`
	NewRecord            string                        `gorm:"type:varchar"`
}

func ToDomain(rec *MedicalRecord) medicalRecordEntity.Domain {
	return medicalRecordEntity.Domain{
		ID:                   rec.ID,
		Uuid:                 rec.Uuid,
		PatientID:            rec.PatientID,
		PatientName:          rec.PatientName,
		PatientAddress:       rec.PatientAddress,
		PatientDob:           rec.PatientDob,
		PatientHeight:        rec.PatientHeight,
		PatientNik:           rec.PatientNik,
		PatientSex:           rec.PatientSex,
		PatientStatusMartial: rec.PatientStatusMartial,
		PatientWeight:        rec.PatientWeight,
		UserID:               rec.UserID,
		Username:             rec.Username,
		UserRole:             rec.UserRole,
		UserSpecialist:       rec.UserSpecialist,
		MedicalStaffID:       rec.MedicalStaffID,
		MedicalStaff:         rec.MedicalStaff.Name,
		Consultation:         rec.Consultation,
		NewRecord:            rec.NewRecord,
	}
}

func FromDomain(medicalRecordDomain *medicalRecordEntity.Domain) *MedicalRecord {
	return &MedicalRecord{
		ID:                   medicalRecordDomain.ID,
		Uuid:                 medicalRecordDomain.Uuid,
		PatientID:            medicalRecordDomain.PatientID,
		PatientName:          medicalRecordDomain.PatientName,
		PatientAddress:       medicalRecordDomain.PatientAddress,
		PatientDob:           medicalRecordDomain.PatientDob,
		PatientHeight:        medicalRecordDomain.PatientHeight,
		PatientNik:           medicalRecordDomain.PatientNik,
		PatientSex:           medicalRecordDomain.PatientSex,
		PatientStatusMartial: medicalRecordDomain.PatientStatusMartial,
		UserID:               medicalRecordDomain.UserID,
		Username:             medicalRecordDomain.Username,
		UserRole:             medicalRecordDomain.UserRole,
		UserSpecialist:       medicalRecordDomain.UserSpecialist,
		MedicalStaffID:       medicalRecordDomain.MedicalStaffID,
		Consultation:         medicalRecordDomain.Consultation,
		NewRecord:            medicalRecordDomain.NewRecord,
	}
}

func toDomainArray(record []MedicalRecord) []medicalRecordEntity.Domain {
	var res []medicalRecordEntity.Domain
	for _, v := range record {
		res = append(res, ToDomain(&v))
	}
	return res
}
