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
	FindByDay(ctx context.Context, day string) (Domain, error)
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	DeleteWorkDay(ctx context.Context, id string) (string, error)
	GetWorkDays(ctx context.Context) (*[]Domain, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewWorkDay(ctx context.Context, data *Domain) (*Domain, error)
	GetByDay(ctx context.Context, nik string) (Domain, error)
	UpdateWorkDay(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetWorkDays(ctx context.Context) (*[]Domain, error)
	DeleteWorkDayByUuid(ctx context.Context, id string) (string, error)
}
