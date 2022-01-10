package recipeEntity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID              uint
	Uuid            uuid.UUID
	MedicineID      uint
	ConsumptionRule string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Service interface {
	CreateNewRecipe(ctx context.Context, data *Domain) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	UpdateById(ctx context.Context, data *Domain, id string) (*Domain, error)
	DeleteRecipe(ctx context.Context, id string) (string, error)
	GetRecipes(ctx context.Context) (*[]Domain, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewRecipe(ctx context.Context, data *Domain) (*Domain, error)
	UpdateRecipe(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetRecipe(ctx context.Context) (*[]Domain, error)
	DeleteRecipeByUuid(ctx context.Context, id string) (string, error)
}
