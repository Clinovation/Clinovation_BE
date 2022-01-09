package queueEntity

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID         uint
	Uuid       uuid.UUID
	PatientID  uint
	ScheduleID uint
	UserID     uint
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	CreateQueue(ctx context.Context, data *Domain, userID string, scheduleId string, patientId string) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	UpdateById(ctx context.Context, data *Domain, userID string, scheduleId string, patientId string, id string) (*Domain, error)
	DeleteQueue(ctx context.Context, id string) (string, error)
	GetDoctorQueues(ctx context.Context) (*[]Domain, error)
	GetNurseQueues(ctx context.Context) (*[]Domain, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewQueue(ctx context.Context, data *Domain) (*Domain, error)
	UpdateQueue(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetDoctorQueues(ctx context.Context) (*[]Domain, error)
	GetNurseQueues(ctx context.Context) (*[]Domain, error)
	DeleteQueueByUuid(ctx context.Context, id string) (string, error)
}
