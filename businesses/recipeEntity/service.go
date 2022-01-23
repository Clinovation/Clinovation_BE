package recipeEntity

import (
	"context"
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicineEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"log"
	"time"
)

type RecipeServices struct {
	RecipesRepository       Repository
	DoctorsRepository       doctorsEntity.Repository
	NursesRepository        nursesEntity.Repository
	PatientRepository       patientEntity.Repository
	MedicalRecordRepository medicalRecordEntity.Repository
	MedicineRepository      medicineEntity.Repository
	jwtAuth                 *auth.ConfigJWT
	ContextTimeout          time.Duration
}

func NewRecipeServices(repoRecipe Repository, repoDoctor doctorsEntity.Repository, repoNurse nursesEntity.Repository,
	repoPatient patientEntity.Repository, repoSMedicalRecord medicalRecordEntity.Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &RecipeServices{
		RecipesRepository:       repoRecipe,
		DoctorsRepository:       repoDoctor,
		NursesRepository:        repoNurse,
		PatientRepository:       repoPatient,
		MedicalRecordRepository: repoSMedicalRecord,
		jwtAuth:                 auth,
		ContextTimeout:          timeout,
	}
}

func (qs *RecipeServices) CreateNewRecipe(ctx context.Context, recipeDomain *Domain, userID string, medicalRecordID string, medicineID string, patientID string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	doctor, errDoctor := qs.DoctorsRepository.GetByUuid(ctx, userID)
	nurse, errNurse := qs.NursesRepository.GetByUuid(ctx, userID)

	if errDoctor != nil && errNurse != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	patient, errPatient := qs.PatientRepository.GetByUuid(ctx, patientID)

	if errPatient != nil {
		return &Domain{}, errors.New("Patient Doen't Exist")
	}

	medicalRecord, err := qs.MedicalRecordRepository.GetByUuid(ctx, medicalRecordID)
	if err != nil {
		return &Domain{}, errors.New("Medical Record Doen't Exist")
	}

	medicine, err := qs.MedicineRepository.GetByUuid(ctx, medicineID)
	if err != nil {
		return &Domain{}, errors.New("Medicine Doen't Exist")
	}

	if doctor.Role != "" && doctor.Role != "approve_waiting_list" {
		recipeDomain.UserID = doctor.ID
		recipeDomain.UserRole = doctor.Role
		recipeDomain.Username = doctor.Name
	}

	if nurse.Role != "" && nurse.Role != "approve_waiting_list" {
		recipeDomain.UserID = nurse.ID
		recipeDomain.UserRole = nurse.Role
		recipeDomain.Username = nurse.Name
	}

	recipeDomain.PatientID = patient.ID
	recipeDomain.MedicalRecordID = medicalRecord.ID
	recipeDomain.MedicineID = medicine.ID

	res, err := qs.RecipesRepository.CreateNewRecipe(ctx, recipeDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (qs *RecipeServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	result, err := qs.RecipesRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (qs *RecipeServices) UpdateById(ctx context.Context, recipeDomain *Domain, id string, userID string, medicalRecordID string, medicineID string, patientID string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	doctor, errDoctor := qs.DoctorsRepository.GetByUuid(ctx, userID)
	nurse, errNurse := qs.NursesRepository.GetByUuid(ctx, userID)

	if errDoctor != nil && errNurse != nil {
		return &Domain{}, errors.New("User Doen't Exist")
	}

	patient, errPatient := qs.PatientRepository.GetByUuid(ctx, patientID)

	if errPatient != nil {
		return &Domain{}, errors.New("Patient Doen't Exist")
	}

	medicalRecord, err := qs.MedicalRecordRepository.GetByUuid(ctx, medicalRecordID)
	if err != nil {
		return &Domain{}, errors.New("Medical Record Doen't Exist")
	}

	medicine, err := qs.MedicineRepository.GetByUuid(ctx, medicineID)
	if err != nil {
		return &Domain{}, errors.New("Medicine Doen't Exist")
	}

	if doctor.Role != "" && doctor.Role != "approve_waiting_list" {
		recipeDomain.UserID = doctor.ID
		recipeDomain.UserRole = doctor.Role
		recipeDomain.Username = doctor.Name
	}

	if nurse.Role != "" && nurse.Role != "approve_waiting_list" {
		recipeDomain.UserID = nurse.ID
		recipeDomain.UserRole = nurse.Role
		recipeDomain.Username = nurse.Name
	}

	recipeDomain.PatientID = patient.ID
	recipeDomain.MedicalRecordID = medicalRecord.ID
	recipeDomain.MedicineID = medicine.ID

	result, err := qs.RecipesRepository.UpdateRecipe(ctx, id, recipeDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (qs *RecipeServices) DeleteRecipe(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	res, err := qs.RecipesRepository.DeleteRecipeByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundRecipe
	}
	return res, nil
}

func (qs *RecipeServices) GetRecipes(ctx context.Context, page int) (*[]Domain, int, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	var offset int
	limit := 5
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 5
	}

	res, totalData, err := qs.RecipesRepository.GetRecipe(ctx, offset, limit)
	if err != nil {
		log.Println(err)
		return &[]Domain{}, -1, -1, -1, businesses.ErrNotFoundRecipe
	}

	return res, offset, limit, totalData, nil
}
