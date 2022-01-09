package medicalStaffEntity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID             uint
	Uuid           uuid.UUID
	Nik            string
	Name           string
	Email          string
	Dob            string
	Sex            string
	Contact        string
	Password       string
	Role           string
	WorkExperience string
	Avatar         string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Service interface {
	Register(ctx context.Context, data *Domain) (*Domain, error)
	Login(ctx context.Context, email string, password string) (string, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	FindByName(ctx context.Context, name string) ([]Domain, error)
	FindByNik(ctx context.Context, name string) ([]Domain, error)
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, fileLocation string) (*Domain, error)
	DeleteMedicalStaff(ctx context.Context, id string) (string, error)
	GetMedicalStaff(ctx context.Context) (*[]Domain, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewMedicalStaff(ctx context.Context, data *Domain) (*Domain, error)
	UpdateMedicalStaff(ctx context.Context, id string, data *Domain) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByNik(ctx context.Context, nik string) (Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetMedicalStaff(ctx context.Context) (*[]Domain, error)
	GetByName(ctx context.Context, name string) ([]Domain, error)
	GetByNikByQuery(ctx context.Context, nik string) ([]Domain, error)
	DeleteMedicalStaffByUuid(ctx context.Context, id string) (string, error)
}
