package scheduleEntity

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID         uint
	Uuid       uuid.UUID
	WorkHourID uint
	WorkDayID  uint
	UserID     uint
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	CreateSchedule(ctx context.Context, data *Domain, userID string, workDayId string, workHourId string) (*Domain, error)
	FindByUuid(ctx context.Context, uuid string) (Domain, error)
	UpdateById(ctx context.Context, data *Domain, userID string, workDayId string, workHourId string, id string) (*Domain, error)
	DeleteSchedule(ctx context.Context, id string) (string, error)
	GetDoctorSchedules(ctx context.Context) (*[]Domain, error)
	GetNurseSchedules(ctx context.Context) (*[]Domain, error)
	GetDoctorSchedulesByDay(ctx context.Context, day string) (*[]Domain, error)
	GetDoctorSchedulesByHour(ctx context.Context, hour string) (*[]Domain, error)
	GetNurseSchedulesByDay(ctx context.Context, day string) (*[]Domain, error)
	GetNurseSchedulesByHour(ctx context.Context, hour string) (*[]Domain, error)
}

type Repository interface {
	// Databases postgresql
	CreateNewSchedule(ctx context.Context, data *Domain) (*Domain, error)
	UpdateSchedule(ctx context.Context, id string, data *Domain) (*Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	GetByDay(ctx context.Context, day string) (*[]Domain, error)
	GetByHour(ctx context.Context, hour string) (*[]Domain, error)
	GetByUuid(ctx context.Context, uuid string) (Domain, error)
	GetDoctorSchedules(ctx context.Context) (*[]Domain, error)
	GetNurseSchedules(ctx context.Context) (*[]Domain, error)
	DeleteScheduleByUuid(ctx context.Context, id string) (string, error)
}
