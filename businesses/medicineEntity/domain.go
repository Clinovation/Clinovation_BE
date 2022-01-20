package medicineEntity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID        uint
	Uuid      uuid.UUID
	Name      string
	Type      string
	Price     int
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	CreateNewMedicine(ctx context.Context, data *Domain) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	FindByName(ctx context.Context, name string, page int) ([]Domain, int, int, int64, error)
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	DeleteMedicine(ctx context.Context, id string) (string, error)
	GetMedicines(ctx context.Context, page int) (*[]Domain, int, int, int64, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewMedicine(ctx context.Context, data *Domain) (*Domain, error)
	UpdateMedicine(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetMedicine(ctx context.Context, offset, limit int) (*[]Domain, int64, error)
	GetByNameByQuery(ctx context.Context, name string, offset, limit int) ([]Domain, int64, error)
	GetByName(ctx context.Context, name string) (Domain, error)
	DeleteMedicineByUuid(ctx context.Context, id string) (string, error)
}
