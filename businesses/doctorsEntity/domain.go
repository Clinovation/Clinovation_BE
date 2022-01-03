package doctorsEntity

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
	Specialist     string
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
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, fileLocation string) (*Domain, error)
	DeleteDoctor(ctx context.Context, id string) (string, error)
	GetDoctors(ctx context.Context) (*[]Domain, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewDoctor(ctx context.Context, data *Domain) (*Domain, error)
	UpdateDoctor(ctx context.Context, id string, data *Domain) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByNik(ctx context.Context, nik string) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetDoctors(ctx context.Context) (*[]Domain, error)
	DeleteDoctorByUuid(ctx context.Context, id string) (string, error)
}
