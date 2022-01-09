package scheduleRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/scheduleEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) scheduleEntity.Repository {
	return &ScheduleRepository{
		db: db,
	}
}

func (r *ScheduleRepository) CreateNewSchedule(ctx context.Context, scheduleDomain *scheduleEntity.Domain) (*scheduleEntity.Domain, error) {
	rec := FromDomain(scheduleDomain)
	rec.Uuid, _ = uuid.NewRandom()
	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *ScheduleRepository) GetByUuid(ctx context.Context, uuid string) (scheduleEntity.Domain, error) {
	rec := Schedule{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return scheduleEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *ScheduleRepository) UpdateSchedule(ctx context.Context, id string, scheduleDomain *scheduleEntity.Domain) (*scheduleEntity.Domain, error) {
	rec := FromDomain(scheduleDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &scheduleEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &scheduleEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *ScheduleRepository) GetByID(ctx context.Context, id uint) (scheduleEntity.Domain, error) {
	rec := Schedule{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return scheduleEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *ScheduleRepository) DeleteScheduleByUuid(ctx context.Context, id string) (string, error) {
	rec := Schedule{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Schedule Data was Deleted", nil
}

func (r *ScheduleRepository) GetDoctorSchedules(ctx context.Context) (*[]scheduleEntity.Domain, error) {
	var schedule []Schedule
	if err := r.db.Find(&schedule).Error; err != nil {
		return &[]scheduleEntity.Domain{}, err
	}
	result := toDomainArray(schedule)
	return &result, nil
}

func (r *ScheduleRepository) GetNurseSchedules(ctx context.Context) (*[]scheduleEntity.Domain, error) {
	var schedule []Schedule
	if err := r.db.Find(&schedule).Error; err != nil {
		return &[]scheduleEntity.Domain{}, err
	}
	result := toDomainArray(schedule)
	return &result, nil
}

func (r *ScheduleRepository) GetByHour(ctx context.Context, hour string) (*[]scheduleEntity.Domain, error) {
	rec := []Schedule{}

	err := r.db.Find(&rec, "hour LIKE ?", "%"+hour+"%").Error
	if err != nil {
		return nil, err
	}
	result := toDomainArray(rec)

	return &result, nil
}

func (r *ScheduleRepository) GetByDay(ctx context.Context, day string) (*[]scheduleEntity.Domain, error) {
	rec := []Schedule{}

	err := r.db.Find(&rec, "day LIKE ?", "%"+day+"%").Error
	if err != nil {
		return nil, err
	}
	result := toDomainArray(rec)

	return &result, nil
}
