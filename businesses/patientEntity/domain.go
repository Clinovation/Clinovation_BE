package patientEntity

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID             uint
	MedicalStaffID uint
	MedicalStaff   string
	Uuid           uuid.UUID
	Nik            string
	Name           string
	Dob            string
	Sex            string
	Contact        string
	StatusMartial  string
	Address        string
	Height         string
	Weight         string
	Role           string
	Avatar         string
	Record         string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Service interface {
	Register(ctx context.Context, data *Domain, medicalStaffID string) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	FindByName(ctx context.Context, name string, page int) ([]Domain, int, int, int64, error)
	FindByNik(ctx context.Context, nik string, page int) ([]Domain, int, int, int64, error)
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, fileLocation string) (*Domain, error)
	DeletePatient(ctx context.Context, id string) (string, error)
	GetPatients(ctx context.Context, page int) (*[]Domain, int, int, int64, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewPatient(ctx context.Context, data *Domain) (*Domain, error)
	UpdatePatient(ctx context.Context, id string, data *Domain) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByNik(ctx context.Context, nik string) (Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetPatients(ctx context.Context, offset, limit int) (*[]Domain, int64, error)
	GetByName(ctx context.Context, name string, offset, limit int) ([]Domain, int64, error)
	GetByNikByQuery(ctx context.Context, nik string, offset, limit int) ([]Domain, int64, error)
	DeletePatientByUuid(ctx context.Context, id string) (string, error)
}
