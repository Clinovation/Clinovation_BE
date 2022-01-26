package nursesEntity

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
	WorkDay        string
	WorkDayID      uint
	WorkHour       string
	WorkHourID     uint
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Service interface {
	Register(ctx context.Context, data *Domain, workDayID string, workHourID string) (*Domain, error)
	Login(ctx context.Context, email string, password string) (string, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	FindByName(ctx context.Context, name string, page int) ([]Domain, int, int, int64, error)
	FindByNik(ctx context.Context, nik string, page int) ([]Domain, int, int, int64, error)
	ForgetPassword(ctx context.Context, data *Domain) (*Domain, error)
	ChangePassword(ctx context.Context, data *Domain, id string) (*Domain, error)
	AcceptNurse(ctx context.Context, id string) (*Domain, error)
	UpdateById(ctx context.Context, data *Domain, id string, workDayID string, workHourID string) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, fileLocation string) (*Domain, error)
	DeleteNurse(ctx context.Context, id string) (string, error)
	GetNurses(ctx context.Context, page int) (*[]Domain, int, int, int64, error)
	GetWaitingList(ctx context.Context, page int) (*[]Domain, int, int, int64, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewNurse(ctx context.Context, data *Domain) (*Domain, error)
	UpdateNurse(ctx context.Context, id string, data *Domain) (*Domain, error)
	AcceptNurse(ctx context.Context, id string, data *Domain) (*Domain, error)
	UploadAvatar(ctx context.Context, id string, data *Domain) (*Domain, error)
	ForgetPassword(ctx context.Context, nik string, email string) (Domain, error)
	GetByNik(ctx context.Context, nik string) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetNurses(ctx context.Context, offset, limit int) (*[]Domain, int64, error)
	GetWaitingList(ctx context.Context, offset, limit int) (*[]Domain, int64, error)
	GetByName(ctx context.Context, name string, offset, limit int) ([]Domain, int64, error)
	GetByNikByQuery(ctx context.Context, nik string, offset, limit int) ([]Domain, int64, error)
	DeleteNurseByUuid(ctx context.Context, id string) (string, error)
}
