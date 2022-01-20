package workDayRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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

func (r *WorkDaysRepository) GetByDay(ctx context.Context, day string) (workDayEntity.Domain, error) {
	rec := WorkDays{}

	err := r.db.Where("day = ?", day).First(&rec).Error
	if err != nil {
		return workDayEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *WorkDaysRepository) GetByDayByQuery(ctx context.Context, day string, offset, limit int) ([]workDayEntity.Domain, int64,error) {
	var totalData int64
	domain := []workDayEntity.Domain{}
	rec := []WorkDays{}

	r.db.Find(&rec, "day LIKE ?", "%"+day+"%").Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Find(&rec, "day LIKE ?", "%"+day+"%").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)

	return domain, totalData, nil
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

func (r *WorkDaysRepository) GetWorkDays(ctx context.Context, offset, limit int) (*[]workDayEntity.Domain,int64, error) {
	var totalData int64
	domain := []workDayEntity.Domain{}
	rec := []WorkDays{}

	r.db.Find(&rec).Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)

	return &domain, totalData, nil
}
