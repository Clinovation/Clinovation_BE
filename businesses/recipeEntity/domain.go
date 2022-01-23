package recipeEntity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID              uint
	Uuid            uuid.UUID
	MedicalRecordID uint
	MedicineID      uint
	Medicine        string
	PatientID       uint
	Patient         string
	UserID          uint
	Username        string
	UserRole        string
	ConsumptionRule string
	Symptom         string
	Record          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Service interface {
	CreateNewRecipe(ctx context.Context, data *Domain, userID string, medicalRecordID string, medicineID string, patientID string) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	UpdateById(ctx context.Context, data *Domain, id string, userID string, medicalRecordID string, medicineID string, patientID string) (*Domain, error)
	DeleteRecipe(ctx context.Context, id string) (string, error)
	GetRecipes(ctx context.Context, page int) (*[]Domain, int, int, int64, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewRecipe(ctx context.Context, data *Domain) (*Domain, error)
	UpdateRecipe(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetRecipe(ctx context.Context, offset, limit int) (*[]Domain, int64, error)
	DeleteRecipeByUuid(ctx context.Context, id string) (string, error)
}
