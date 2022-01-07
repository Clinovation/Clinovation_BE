package workHourRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/workHourEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkHoursRepository struct {
	db *gorm.DB
}

func NewWorkHoursRepository(db *gorm.DB) workHourEntity.Repository {
	return &WorkHoursRepository{
		db: db,
	}
}

func (r *WorkHoursRepository) CreateNewWorkHour(ctx context.Context, workHourDomain *workHourEntity.Domain) (*workHourEntity.Domain, error) {
	rec := FromDomain(workHourDomain)
	rec.Uuid, _ = uuid.NewRandom()

	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *WorkHoursRepository) GetByHour(ctx context.Context, nik string) (workHourEntity.Domain, error) {
	rec := WorkHours{}

	err := r.db.Where("hour = ?", nik).First(&rec).Error
	if err != nil {
		return workHourEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *WorkHoursRepository) GetByUuid(ctx context.Context, uuid string) (workHourEntity.Domain, error) {
	rec := WorkHours{}

	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return workHourEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *WorkHoursRepository) UpdateWorkHour(ctx context.Context, id string, workHourDomain *workHourEntity.Domain) (*workHourEntity.Domain, error) {
	rec := FromDomain(workHourDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &workHourEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &workHourEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *WorkHoursRepository) DeleteWorkHourByUuid(ctx context.Context, id string) (string, error) {
	rec := WorkHours{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Work Hour Was Deleted", nil
}

func (r *WorkHoursRepository) GetWorkHours(ctx context.Context) (*[]workHourEntity.Domain, error) {
	var workHour []WorkHours

	if err := r.db.Find(&workHour).Error; err != nil {
		return &[]workHourEntity.Domain{}, err
	}
	result := toDomainArray(workHour)
	return &result, nil
}
