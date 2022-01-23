package medicalRecordEntity

import (
	"context"
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"log"
	"time"
)

type MedicalRecordServices struct {
	MedicalRecordsRepository Repository
	DoctorsRepository        doctorsEntity.Repository
	NursesRepository         nursesEntity.Repository
	PatientRepository        patientEntity.Repository
	MedicalStaffRepository   medicalStaffEntity.Repository
	jwtAuth                  *auth.ConfigJWT
	ContextTimeout           time.Duration
}

func NewMedicalRecordServices(repoMedicalRecord Repository, repoDoctor doctorsEntity.Repository, repoNurse nursesEntity.Repository,
	repoSMedicalStaff medicalStaffEntity.Repository, repoPatient patientEntity.Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &MedicalRecordServices{
		MedicalRecordsRepository: repoMedicalRecord,
		DoctorsRepository:        repoDoctor,
		NursesRepository:         repoNurse,
		MedicalStaffRepository:   repoSMedicalStaff,
		PatientRepository:        repoPatient,
		jwtAuth:                  auth,
		ContextTimeout:           timeout,
	}
}

func (ss *MedicalRecordServices) CreateMedicalRecord(ctx context.Context, medicalRecordDomain *Domain, userID string, medicalStaffID string, patientId string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	doctor, errDoctor := ss.DoctorsRepository.GetByUuid(ctx, userID)
	nurse, errNurse := ss.NursesRepository.GetByUuid(ctx, userID)

	if errDoctor != nil && errNurse != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	medicalStaff, err := ss.MedicalStaffRepository.GetByUuid(ctx, medicalStaffID)
	if err != nil {
		return &Domain{}, errors.New("Medical Staff Doen't Exist")
	}

	patient, err := ss.PatientRepository.GetByUuid(ctx, patientId)
	if err != nil {
		return &Domain{}, errors.New("Data Patient Doen't Exist")
	}

	if doctor.Role != "" && doctor.Role != "approve_waiting_list" {
		medicalRecordDomain.UserID = doctor.ID
		medicalRecordDomain.UserRole = doctor.Role
		medicalRecordDomain.Username = doctor.Name
		medicalRecordDomain.UserSpecialist = doctor.Specialist
	}

	if nurse.Role != "" && nurse.Role != "approve_waiting_list" {
		medicalRecordDomain.UserID = nurse.ID
		medicalRecordDomain.UserRole = nurse.Role
		medicalRecordDomain.Username = nurse.Name
		medicalRecordDomain.UserSpecialist = ""
	}

	medicalRecordDomain.MedicalStaffID = medicalStaff.ID
	medicalRecordDomain.PatientID = patient.ID


	log.Println("service",medicalRecordDomain.NewRecord)
	log.Println("service",medicalRecordDomain.Consultation)

	res, err := ss.MedicalRecordsRepository.CreateNewMedicalRecord(ctx, medicalRecordDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	if medicalRecordDomain.NewRecord != "" {
		patient.Record = patient.Record + "," + medicalRecordDomain.NewRecord
		_, errNewRecord := ss.PatientRepository.UpdatePatient(ctx, patient.Uuid.String(), &patient)
		if errNewRecord != nil {
			return nil, businesses.ErrInternalServer
		}
	}

	return res, nil
}

func (ss *MedicalRecordServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	result, err := ss.MedicalRecordsRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (ss *MedicalRecordServices) UpdateById(ctx context.Context, medicalRecordDomain *Domain, userID string, medicalStaffID string, patientId string, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	_, err := ss.MedicalRecordsRepository.GetByUuid(ctx, id)
	if err != nil {
		return &Domain{}, errors.New("Medical Record Doesn't Exist")
	}

	doctor, errDoctor := ss.DoctorsRepository.GetByUuid(ctx, userID)
	nurse, errNurse := ss.NursesRepository.GetByUuid(ctx, userID)

	if errDoctor != nil && errNurse != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	medicalStaff, err := ss.MedicalStaffRepository.GetByUuid(ctx, medicalStaffID)
	if err != nil {
		return &Domain{}, errors.New("Medical Staff Doen't Exist")
	}

	patient, err := ss.PatientRepository.GetByUuid(ctx, patientId)
	if err != nil {
		return &Domain{}, errors.New("Data Patient Doen't Exist")
	}

	if doctor.Role != "" && doctor.Role != "approve_waiting_list" {
		medicalRecordDomain.UserID = doctor.ID
		medicalRecordDomain.UserRole = doctor.Role
		medicalRecordDomain.Username = doctor.Name
		medicalRecordDomain.UserSpecialist = doctor.Specialist
	}

	if nurse.Role != "" && nurse.Role != "approve_waiting_list" {
		medicalRecordDomain.UserID = nurse.ID
		medicalRecordDomain.UserRole = nurse.Role
		medicalRecordDomain.Username = nurse.Name
		medicalRecordDomain.UserSpecialist = ""
	}

	medicalRecordDomain.MedicalStaffID = medicalStaff.ID
	medicalRecordDomain.PatientID = patient.ID

	res, err := ss.MedicalRecordsRepository.UpdateMedicalRecord(ctx, id, medicalRecordDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (ss *MedicalRecordServices) DeleteMedicalRecord(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	res, err := ss.MedicalRecordsRepository.DeleteMedicalRecordByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundMedicalRecord
	}
	return res, nil
}

func (ss *MedicalRecordServices) GetMedicalRecordsQueue(ctx context.Context, userID string, page int) (*[]Domain, int, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	var userId uint

	doctor, errDoctor := ss.DoctorsRepository.GetByUuid(ctx, userID)
	nurse, errNurse := ss.NursesRepository.GetByUuid(ctx, userID)

	if errDoctor != nil && errNurse != nil {
		return &[]Domain{}, -1, -1, -1, errors.New("User Doen't Exist")
	}

	if doctor.ID != 0{
		userId = doctor.ID
	}else {
		userId = nurse.ID
	}

	var offset int
	limit := 5
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 5
	}

	res, totalData, err := ss.MedicalRecordsRepository.GetMedicalRecordsQueue(ctx, userId, offset, limit)
	if err != nil {
		log.Println(err)
		return &[]Domain{}, -1, -1, -1, businesses.ErrNotFoundMedicalRecord
	}

	return res, offset, limit, totalData, nil
}
