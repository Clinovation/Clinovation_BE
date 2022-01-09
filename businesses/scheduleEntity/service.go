package scheduleEntity

import (
	"context"
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/workHourEntity"
	"time"
)

type ScheduleServices struct {
	SchedulesRepository Repository
	DoctorsRepository   doctorsEntity.Repository
	NursesRepository    nursesEntity.Repository
	WorkDaysRepository  workDayEntity.Repository
	WorkHoursRepository workHourEntity.Repository
	jwtAuth             *auth.ConfigJWT
	ContextTimeout      time.Duration
}

func NewScheduleServices(repoSchedule Repository, repoDoctor doctorsEntity.Repository, repoNurse nursesEntity.Repository,
	repoWorkDay workDayEntity.Repository, repoWorkHour workHourEntity.Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &ScheduleServices{
		SchedulesRepository: repoSchedule,
		DoctorsRepository:   repoDoctor,
		NursesRepository:    repoNurse,
		WorkDaysRepository:  repoWorkDay,
		WorkHoursRepository: repoWorkHour,
		jwtAuth:             auth,
		ContextTimeout:      timeout,
	}
}

func (ss *ScheduleServices) CreateSchedule(ctx context.Context, scheduleDomain *Domain, userID string, workDayId string, workHourId string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	user, err := ss.SchedulesRepository.GetByUuid(ctx, userID)
	if err != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	workDay, err := ss.WorkDaysRepository.GetByUuid(ctx, workDayId)
	if err != nil {
		return &Domain{}, errors.New("Work Day Doen't Exist")
	}

	workHour, err := ss.WorkHoursRepository.GetByUuid(ctx, workHourId)
	if err != nil {
		return &Domain{}, errors.New("Work Hour Doen't Exist")
	}

	scheduleDomain.UserID = user.ID
	scheduleDomain.Role = user.Role
	scheduleDomain.WorkHourID = workHour.ID
	scheduleDomain.WorkDayID = workDay.ID

	res, err := ss.SchedulesRepository.CreateNewSchedule(ctx, scheduleDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (ss *ScheduleServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	result, err := ss.SchedulesRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (ss *ScheduleServices) UpdateById(ctx context.Context, scheduleDomain *Domain, userID string, workDayId string, workHourId string, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	user, err := ss.SchedulesRepository.GetByUuid(ctx, userID)
	if err != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	workDay, err := ss.WorkDaysRepository.GetByUuid(ctx, workDayId)
	if err != nil {
		return &Domain{}, errors.New("Work Day Doen't Exist")
	}

	workHour, err := ss.WorkHoursRepository.GetByUuid(ctx, workHourId)
	if err != nil {
		return &Domain{}, errors.New("Work Hour Doen't Exist")
	}

	scheduleDomain.UserID = user.ID
	scheduleDomain.Role = user.Role
	scheduleDomain.WorkHourID = workHour.ID
	scheduleDomain.WorkDayID = workDay.ID

	res, err := ss.SchedulesRepository.UpdateSchedule(ctx, id, scheduleDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (ss *ScheduleServices) DeleteSchedule(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	res, err := ss.SchedulesRepository.DeleteScheduleByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundSchedule
	}
	return res, nil
}

func (ss *ScheduleServices) GetDoctorSchedules(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	res, err := ss.SchedulesRepository.GetDoctorSchedules(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}

func (ss *ScheduleServices) GetNurseSchedules(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	res, err := ss.SchedulesRepository.GetNurseSchedules(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}

func (ss *ScheduleServices) GetDoctorSchedulesByDay(ctx context.Context, day string) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	result, err := ss.SchedulesRepository.GetByDay(ctx, day)
	if err != nil {
		return &[]Domain{}, err
	}

	return result, nil
}

func (ss *ScheduleServices) GetDoctorSchedulesByHour(ctx context.Context, hour string) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	result, err := ss.SchedulesRepository.GetByDay(ctx, hour)
	if err != nil {
		return &[]Domain{}, err
	}

	return result, nil
}

func (ss *ScheduleServices) GetNurseSchedulesByHour(ctx context.Context, hour string) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	result, err := ss.SchedulesRepository.GetByDay(ctx, hour)
	if err != nil {
		return &[]Domain{}, err
	}

	return result, nil
}

func (ss *ScheduleServices) GetNurseSchedulesByDay(ctx context.Context, day string) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	result, err := ss.SchedulesRepository.GetByDay(ctx, day)
	if err != nil {
		return &[]Domain{}, err
	}

	return result, nil
}
