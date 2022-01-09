package workDayRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkDaysRepository struct {
	db *gorm.DB
}

func NewWorkDaysRepository(db *gorm.DB) workDayEntity.Repository {
	return &WorkDaysRepository{
		db: db,
	}
}

func (r *WorkDaysRepository) CreateNewWorkDay(ctx context.Context, workDayDomain *workDayEntity.Domain) (*workDayEntity.Domain, error) {
	rec := FromDomain(workDayDomain)
	rec.Uuid, _ = uuid.NewRandom()

	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *WorkDaysRepository) GetByDay(ctx context.Context, nik string) (workDayEntity.Domain, error) {
	rec := WorkDays{}

	err := r.db.Where("day = ?", nik).First(&rec).Error
	if err != nil {
		return workDayEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *WorkDaysRepository) GetByUuid(ctx context.Context, uuid string) (workDayEntity.Domain, error) {
	rec := WorkDays{}

	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return workDayEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *WorkDaysRepository) GetByID(ctx context.Context, id uint) (workDayEntity.Domain, error) {
	rec := WorkDays{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return workDayEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *WorkDaysRepository) UpdateWorkDay(ctx context.Context, id string, workDayDomain *workDayEntity.Domain) (*workDayEntity.Domain, error) {
	rec := FromDomain(workDayDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &workDayEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &workDayEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *WorkDaysRepository) DeleteWorkDayByUuid(ctx context.Context, id string) (string, error) {
	rec := WorkDays{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Work Day Was Deleted", nil
}

func (r *WorkDaysRepository) GetWorkDays(ctx context.Context) (*[]workDayEntity.Domain, error) {
	var workDay []WorkDays

	if err := r.db.Find(&workDay).Error; err != nil {
		return &[]workDayEntity.Domain{}, err
	}
	result := toDomainArray(workDay)
	return &result, nil
}
