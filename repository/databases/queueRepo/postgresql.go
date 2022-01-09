package queueRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/queueEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QueueRepository struct {
	db *gorm.DB
}

func NewQueueRepository(db *gorm.DB) queueEntity.Repository {
	return &QueueRepository{
		db: db,
	}
}

func (r *QueueRepository) CreateNewQueue(ctx context.Context, queueDomain *queueEntity.Domain) (*queueEntity.Domain, error) {
	rec := FromDomain(queueDomain)
	rec.Uuid, _ = uuid.NewRandom()
	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *QueueRepository) GetByUuid(ctx context.Context, uuid string) (queueEntity.Domain, error) {
	rec := Queue{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return queueEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *QueueRepository) UpdateQueue(ctx context.Context, id string, queueDomain *queueEntity.Domain) (*queueEntity.Domain, error) {
	rec := FromDomain(queueDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &queueEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &queueEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *QueueRepository) GetByID(ctx context.Context, id uint) (queueEntity.Domain, error) {
	rec := Queue{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return queueEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *QueueRepository) DeleteQueueByUuid(ctx context.Context, id string) (string, error) {
	rec := Queue{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Queue Data was Deleted", nil
}

func (r *QueueRepository) GetDoctorQueues(ctx context.Context) (*[]queueEntity.Domain, error) {
	var queue []Queue
	if err := r.db.Find(&queue).Error; err != nil {
		return &[]queueEntity.Domain{}, err
	}
	result := toDomainArray(queue)
	return &result, nil
}

func (r *QueueRepository) GetNurseQueues(ctx context.Context) (*[]queueEntity.Domain, error) {
	var queue []Queue
	if err := r.db.Find(&queue).Error; err != nil {
		return &[]queueEntity.Domain{}, err
	}
	result := toDomainArray(queue)
	return &result, nil
}
