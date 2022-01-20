package workHourEntity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID        uint
	Uuid      uuid.UUID
	Hour      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	CreateWorkHour(ctx context.Context, data *Domain) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	FindByHour(ctx context.Context, hour string, page int) ([]Domain, int, int, int64, error)
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	DeleteWorkHour(ctx context.Context, id string) (string, error)
	GetWorkHours(ctx context.Context, page int) (*[]Domain, int, int, int64, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewWorkHour(ctx context.Context, data *Domain) (*Domain, error)
	GetByHour(ctx context.Context, hour string) (Domain, error)
	GetByHourByQuery(ctx context.Context, hour string, offset, limit int) ([]Domain, int64, error)
	UpdateWorkHour(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetWorkHours(ctx context.Context, offset, limit int) (*[]Domain, int64, error)
	DeleteWorkHourByUuid(ctx context.Context, id string) (string, error)
}
