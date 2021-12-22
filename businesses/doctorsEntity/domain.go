package doctorsEntity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Domain struct {
	ID          uint
	Uuid        uuid.UUID
	Name        string
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
	Avatar      string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	Register(ctx context.Context, data *Domain) (*Domain, error)
	Login(ctx context.Context, email string, password string) (string, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	Logout(ctx echo.Context) error
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, fileLocation string) (*Domain, error)
	DeleteDoctor(ctx context.Context, id string) (string, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewDoctor(ctx context.Context, data *Domain) (*Domain, error)
	UpdateDoctor(ctx context.Context, id string, data *Domain) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	DeleteDoctorByUuid(ctx context.Context, id string) (string, error)
}
