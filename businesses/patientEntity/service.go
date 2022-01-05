package patientEntity

import (
	"context"
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"strings"
	"time"
)

type PatientServices struct {
	PatientsRepository     Repository
	MedicalStaffRepository medicalStaffEntity.Repository
	jwtAuth                *auth.ConfigJWT
	ContextTimeout         time.Duration
}

func NewPatientServices(repoPatient Repository, repoMedicalStaff medicalStaffEntity.Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &PatientServices{
		PatientsRepository:     repoPatient,
		MedicalStaffRepository: repoMedicalStaff,
		jwtAuth:                auth,
		ContextTimeout:         timeout,
	}
}

func (ps *PatientServices) Register(ctx context.Context, patientDomain *Domain, medicalStaffID string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	existedPatientByNik, err := ps.PatientsRepository.GetByNik(ctx, patientDomain.Nik)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedPatientByNik != (Domain{}) {
		return nil, businesses.ErrDuplicateNik
	}

	medicalStaff, err := ps.MedicalStaffRepository.GetByUuid(ctx, medicalStaffID)
	if err != nil {
		return &Domain{}, errors.New("Medical is not login")
	}

	patientDomain.MedicalStaffID = medicalStaff.ID

	res, err := ps.PatientsRepository.CreateNewPatient(ctx, patientDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (ps *PatientServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	result, err := ps.PatientsRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (ps *PatientServices) UpdateById(ctx context.Context, patientDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	result, err := ps.PatientsRepository.UpdatePatient(ctx, id, patientDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (ps *PatientServices) UploadAvatar(ctx context.Context, id string, imageLink string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	patient, err := ps.PatientsRepository.GetByUuid(ctx, id)
	if err != nil {
		return &Domain{}, err
	}

	patient.Avatar = imageLink
	updateAvatar, err := ps.PatientsRepository.UploadAvatar(ctx, id, &patient)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}

func (ps *PatientServices) DeletePatient(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	res, err := ps.PatientsRepository.DeletePatientByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundPatient
	}
	return res, nil
}

func (ps *PatientServices) GetPatients(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	res, err := ps.PatientsRepository.GetPatients(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}

func (ps *PatientServices) FindByName(ctx context.Context, name string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	res, err := ps.PatientsRepository.GetByName(ctx, name)
	if err != nil {
		return []Domain{}, businesses.ErrNotFoundPatient
	}

	return res, nil
}

func (ps *PatientServices) FindByNik(ctx context.Context, nik string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	res, err := ps.PatientsRepository.GetByNikByQuery(ctx, nik)
	if err != nil {
		return []Domain{}, businesses.ErrNotFoundPatient
	}

	return res, nil
}
