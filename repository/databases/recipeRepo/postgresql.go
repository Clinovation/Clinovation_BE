package recipeRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/recipeEntity"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
)

type RecipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository(db *gorm.DB) recipeEntity.Repository {
	return &RecipeRepository{
		db: db,
	}
}

func (r *RecipeRepository) CreateNewRecipe(ctx context.Context, recipeDomain *recipeEntity.Domain) (*recipeEntity.Domain, error) {
	rec := FromDomain(recipeDomain)
	rec.Uuid, _ = uuid.NewRandom()
	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *RecipeRepository) GetByID(ctx context.Context, id uint) (recipeEntity.Domain, error) {
	rec := Recipe{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return recipeEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *RecipeRepository) GetByUuid(ctx context.Context, uuid string) (recipeEntity.Domain, error) {
	rec := Recipe{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return recipeEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *RecipeRepository) UpdateRecipe(ctx context.Context, id string, recipeDomain *recipeEntity.Domain) (*recipeEntity.Domain, error) {
	rec := FromDomain(recipeDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &recipeEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &recipeEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *RecipeRepository) DeleteRecipeByUuid(ctx context.Context, id string) (string, error) {
	rec := Recipe{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Medical Staff was Deleted", nil
}

func (r *RecipeRepository) GetRecipe(ctx context.Context, offset, limit int) (*[]recipeEntity.Domain, int64, error) {
	var totalData int64
	domain := []recipeEntity.Domain{}
	rec := []Recipe{}

	r.db.Find(&rec).Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Joins("Medicine").Find(&rec).Error
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].Medicine = rec[i].Medicine.Name
		domain[i].Patient = rec[i].Patient.Name
	}
	return &domain, totalData, nil
}
