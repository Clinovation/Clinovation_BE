package queueEntity

import (
	"context"
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/scheduleEntity"
	"time"
)

type QueueServices struct {
	QueuesRepository   Repository
	DoctorsRepository  doctorsEntity.Repository
	NursesRepository   nursesEntity.Repository
	PatientRepository  patientEntity.Repository
	ScheduleRepository scheduleEntity.Repository
	jwtAuth            *auth.ConfigJWT
	ContextTimeout     time.Duration
}

func NewQueueServices(repoQueue Repository, repoDoctor doctorsEntity.Repository, repoNurse nursesEntity.Repository,
	repoSchedule scheduleEntity.Repository, repoPatient patientEntity.Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &QueueServices{
		QueuesRepository:   repoQueue,
		DoctorsRepository:  repoDoctor,
		NursesRepository:   repoNurse,
		ScheduleRepository: repoSchedule,
		PatientRepository:  repoPatient,
		jwtAuth:            auth,
		ContextTimeout:     timeout,
	}
}

func (ss *QueueServices) CreateQueue(ctx context.Context, queueDomain *Domain, userID string, scheduleId string, patientId string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	doctor, errDoctor := ss.QueuesRepository.GetByUuid(ctx, userID)
	nurse, errNurse := ss.QueuesRepository.GetByUuid(ctx, userID)

	if errDoctor != nil && errNurse != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	schedule, err := ss.ScheduleRepository.GetByUuid(ctx, scheduleId)
	if err != nil {
		return &Domain{}, errors.New("schedule Doen't Exist")
	}

	patient, err := ss.PatientRepository.GetByUuid(ctx, patientId)
	if err != nil {
		return &Domain{}, errors.New("Data Patient Doen't Exist")
	}

	if doctor.Role != "" && doctor.Role != "approve_waiting_list" {
		queueDomain.Role = "doctor"
		queueDomain.UserID = doctor.ID
	}

	if nurse.Role != "" && nurse.Role != "approve_waiting_list" {
		queueDomain.Role = "nurse"
		queueDomain.UserID = nurse.ID
	}

	queueDomain.ScheduleID = schedule.ID
	queueDomain.PatientID = patient.ID

	res, err := ss.QueuesRepository.CreateNewQueue(ctx, queueDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (ss *QueueServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	result, err := ss.QueuesRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (ss *QueueServices) UpdateById(ctx context.Context, queueDomain *Domain, userID string, scheduleId string, patientId string, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	_, err := ss.QueuesRepository.GetByUuid(ctx, id)
	if err != nil {
		return &Domain{}, errors.New("Queue Doen't Exist")
	}

	doctor, errDoctor := ss.DoctorsRepository.GetByUuid(ctx, userID)
	nurse, errNurse := ss.NursesRepository.GetByUuid(ctx, userID)

	if errDoctor != nil && errNurse != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	schedule, err := ss.ScheduleRepository.GetByUuid(ctx, scheduleId)
	if err != nil {
		return &Domain{}, errors.New("schedule Doen't Exist")
	}

	patient, err := ss.PatientRepository.GetByUuid(ctx, patientId)
	if err != nil {
		return &Domain{}, errors.New("Data Patient Doen't Exist")
	}

	if doctor.Role != "" && doctor.Role != "approve_waiting_list" {
		queueDomain.Role = "doctor"
		queueDomain.UserID = doctor.ID
	}

	if nurse.Role != "" && nurse.Role != "approve_waiting_list" {
		queueDomain.Role = "nurse"
		queueDomain.UserID = nurse.ID
	}

	queueDomain.ScheduleID = schedule.ID
	queueDomain.PatientID = patient.ID

	res, err := ss.QueuesRepository.UpdateQueue(ctx, id, queueDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (ss *QueueServices) DeleteQueue(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	res, err := ss.QueuesRepository.DeleteQueueByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundQueue
	}
	return res, nil
}

func (ss *QueueServices) GetDoctorQueues(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	res, err := ss.QueuesRepository.GetDoctorQueues(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}

func (ss *QueueServices) GetNurseQueues(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	res, err := ss.QueuesRepository.GetNurseQueues(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}
