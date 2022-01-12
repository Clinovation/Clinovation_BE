package medicalRecordEntity

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID           uint
	Uuid         uuid.UUID
	PatientID    uint
	RecipeID     uint
	UserID       uint
	Consultation string
	Note         string
	NewNote      string
	Symptom      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	CreateMedicalRecord(ctx context.Context, data *Domain, userID string, recipeID string, patientID string) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	UpdateById(ctx context.Context, data *Domain, userID string, recipeID string, patientId string, id string) (*Domain, error)
	DeleteMedicalRecord(ctx context.Context, id string) (string, error)
	GetMedicalRecords(ctx context.Context) (*[]Domain, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewMedicalRecord(ctx context.Context, data *Domain) (*Domain, error)
	UpdateMedicalRecord(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetMedicalRecords(ctx context.Context) (*[]Domain, error)
	DeleteMedicalRecordByUuid(ctx context.Context, id string) (string, error)
}
