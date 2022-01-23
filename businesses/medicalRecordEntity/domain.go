package medicalRecordEntity

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID                   uint
	Uuid                 uuid.UUID
	PatientID            uint
	PatientName          string
	PatientAddress       string
	PatientDob           string
	PatientHeight        string
	PatientWeight        string
	PatientNik           string
	PatientSex           string
	PatientStatusMartial string
	UserID               uint
	Username             string
	UserRole             string
	UserSpecialist       string
	MedicalStaffID       uint
	MedicalStaff         string
	Consultation         string
	NewRecord            string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type Service interface {
	CreateMedicalRecord(ctx context.Context, data *Domain, userID string, medicalStaffID string, patientID string) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	UpdateById(ctx context.Context, data *Domain, userID string, medicalStaffID string, patientId string, id string) (*Domain, error)
	DeleteMedicalRecord(ctx context.Context, id string) (string, error)
	GetMedicalRecordsQueue(ctx context.Context, userID string, page int) (*[]Domain, int, int, int64, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewMedicalRecord(ctx context.Context, data *Domain) (*Domain, error)
	UpdateMedicalRecord(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetMedicalRecordsQueue(ctx context.Context, userID uint, offset, limit int) (*[]Domain, int64, error)
	DeleteMedicalRecordByUuid(ctx context.Context, id string) (string, error)
}
