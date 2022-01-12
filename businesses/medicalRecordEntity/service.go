package medicalRecordEntity

import (
	"context"
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/recipeEntity"
	"time"
)

type MedicalRecordServices struct {
	MedicalRecordsRepository Repository
	DoctorsRepository        doctorsEntity.Repository
	NursesRepository         nursesEntity.Repository
	PatientRepository        patientEntity.Repository
	RecipeRepository         recipeEntity.Repository
	jwtAuth                  *auth.ConfigJWT
	ContextTimeout           time.Duration
}

func NewMedicalRecordServices(repoMedicalRecord Repository, repoDoctor doctorsEntity.Repository, repoNurse nursesEntity.Repository,
	repoSRecipe recipeEntity.Repository, repoPatient patientEntity.Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &MedicalRecordServices{
		MedicalRecordsRepository: repoMedicalRecord,
		DoctorsRepository:        repoDoctor,
		NursesRepository:         repoNurse,
		RecipeRepository:         repoSRecipe,
		PatientRepository:        repoPatient,
		jwtAuth:                  auth,
		ContextTimeout:           timeout,
	}
}

func (ss *MedicalRecordServices) CreateMedicalRecord(ctx context.Context, medicalRecordDomain *Domain, userID string, recipeID string, patientId string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	doctor, errDoctor := ss.DoctorsRepository.GetByUuid(ctx, userID)
	nurse, errNurse := ss.NursesRepository.GetByUuid(ctx, userID)

	if errDoctor != nil && errNurse != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	recipe, err := ss.RecipeRepository.GetByUuid(ctx, recipeID)
	if err != nil {
		return &Domain{}, errors.New("Recipe Doen't Exist")
	}

	patient, err := ss.PatientRepository.GetByUuid(ctx, patientId)
	if err != nil {
		return &Domain{}, errors.New("Data Patient Doen't Exist")
	}

	if doctor.Role != "" && doctor.Role != "approve_waiting_list" {
		medicalRecordDomain.UserID = doctor.ID
	}

	if nurse.Role != "" && nurse.Role != "approve_waiting_list" {
		medicalRecordDomain.UserID = nurse.ID
	}

	medicalRecordDomain.RecipeID = recipe.ID
	medicalRecordDomain.PatientID = patient.ID

	if medicalRecordDomain.NewNote != "" {
		medicalRecordDomain.Note = medicalRecordDomain.Note + "," + medicalRecordDomain.NewNote
	}

	res, err := ss.MedicalRecordsRepository.CreateNewMedicalRecord(ctx, medicalRecordDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
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

func (ss *MedicalRecordServices) UpdateById(ctx context.Context, medicalRecordDomain *Domain, userID string, recipeId string, patientId string, id string) (*Domain, error) {
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

	recipe, err := ss.RecipeRepository.GetByUuid(ctx, recipeId)
	if err != nil {
		return &Domain{}, errors.New("Recipe Doen't Exist")
	}

	patient, err := ss.PatientRepository.GetByUuid(ctx, patientId)
	if err != nil {
		return &Domain{}, errors.New("Data Patient Doen't Exist")
	}

	if doctor.Role != "" && doctor.Role != "approve_waiting_list" {
		medicalRecordDomain.UserID = doctor.ID
	}

	if nurse.Role != "" && nurse.Role != "approve_waiting_list" {
		medicalRecordDomain.UserID = nurse.ID
	}

	medicalRecordDomain.RecipeID = recipe.ID
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

func (ss *MedicalRecordServices) GetMedicalRecords(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.ContextTimeout)
	defer cancel()

	res, err := ss.MedicalRecordsRepository.GetMedicalRecords(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}
