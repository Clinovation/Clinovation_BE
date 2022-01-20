package workDayEntity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID        uint
	Uuid      uuid.UUID
	Day       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	CreateWorkDay(ctx context.Context, data *Domain) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	FindByDay(ctx context.Context, day string, page int) ([]Domain, int, int, int64, error)
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	DeleteWorkDay(ctx context.Context, id string) (string, error)
	GetWorkDays(ctx context.Context, page int) (*[]Domain, int, int, int64, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewWorkDay(ctx context.Context, data *Domain) (*Domain, error)
	GetByDay(ctx context.Context, day string) (Domain, error)
	GetByDayByQuery(ctx context.Context, day string, offset, limit int) ([]Domain, int64, error)
	UpdateWorkDay(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetWorkDays(ctx context.Context, offset, limit int) (*[]Domain, int64, error)
	DeleteWorkDayByUuid(ctx context.Context, id string) (string, error)
}
