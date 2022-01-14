package medicalStaffEntity

import (
	"context"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"strings"
	"time"
)

type MedicalStaffServices struct {
	MedicalStaffsRepository Repository
	jwtAuth                 *auth.ConfigJWT
	ContextTimeout          time.Duration
}

func NewMedicalStaffServices(repoMedicalStaff Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &MedicalStaffServices{
		MedicalStaffsRepository: repoMedicalStaff,
		jwtAuth:                 auth,
		ContextTimeout:          timeout,
	}
}

func (mss *MedicalStaffServices) Register(ctx context.Context, medicalStaffDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	existedMedicalStaffByNik, err := mss.MedicalStaffsRepository.GetByNik(ctx, medicalStaffDomain.Nik)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedMedicalStaffByNik != (Domain{}) {
		return nil, businesses.ErrDuplicateNik
	}

	existedMedicalStaffByEmail, err := mss.MedicalStaffsRepository.GetByEmail(ctx, medicalStaffDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedMedicalStaffByEmail != (Domain{}) {
		return nil, businesses.ErrDuplicateEmail
	}

	medicalStaffDomain.Password, err = helpers.HashPassword(medicalStaffDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := mss.MedicalStaffsRepository.CreateNewMedicalStaff(ctx, medicalStaffDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (mss *MedicalStaffServices) Login(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	medicalStaffDomain, err := mss.MedicalStaffsRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", businesses.ErrEmailNotRegistered
	}

	if !helpers.ValidateHash(password, medicalStaffDomain.Password) {
		return "", businesses.ErrPassword
	}

	token := mss.jwtAuth.GenerateToken(medicalStaffDomain.Uuid.String(), medicalStaffDomain.Role)

	return token, nil
}

func (ds *MedicalStaffServices) ChangePassword(ctx context.Context, medicalStaffDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	passwordHash, err := helpers.HashPassword(medicalStaffDomain.Password)
	if err != nil {
		panic(err)
	}

	medicalStaffDomain.Password = passwordHash
	result, err := ds.MedicalStaffsRepository.UpdateMedicalStaff(ctx, id, medicalStaffDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (ds *MedicalStaffServices) ForgetPassword(ctx context.Context, medicalStaffDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	result, err := ds.MedicalStaffsRepository.ForgetPassword(ctx, medicalStaffDomain.Nik, medicalStaffDomain.Email)
	if err != nil {
		return &Domain{}, err
	}
	return &result, nil
}

func (mss *MedicalStaffServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	result, err := mss.MedicalStaffsRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (mss *MedicalStaffServices) UpdateById(ctx context.Context, medicalStaffDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	passwordHash, err := helpers.HashPassword(medicalStaffDomain.Password)
	if err != nil {
		panic(err)
	}

	medicalStaffDomain.Password = passwordHash
	result, err := mss.MedicalStaffsRepository.UpdateMedicalStaff(ctx, id, medicalStaffDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (mss *MedicalStaffServices) UploadAvatar(ctx context.Context, id string, imageLink string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	medicalStaff, err := mss.MedicalStaffsRepository.GetByUuid(ctx, id)
	if err != nil {
		return &Domain{}, err
	}

	medicalStaff.Avatar = imageLink
	updateAvatar, err := mss.MedicalStaffsRepository.UploadAvatar(ctx, id, &medicalStaff)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}

func (mss *MedicalStaffServices) DeleteMedicalStaff(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	res, err := mss.MedicalStaffsRepository.DeleteMedicalStaffByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundMedicalStaff
	}
	return res, nil
}

func (mss *MedicalStaffServices) GetMedicalStaff(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	res, err := mss.MedicalStaffsRepository.GetMedicalStaff(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}

func (mss *MedicalStaffServices) FindByName(ctx context.Context, name string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	res, err := mss.MedicalStaffsRepository.GetByName(ctx, name)
	if err != nil {
		return []Domain{}, businesses.ErrNotFoundMedicalStaff
	}

	return res, nil
}

func (mss *MedicalStaffServices) FindByNik(ctx context.Context, nik string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, mss.ContextTimeout)
	defer cancel()

	res, err := mss.MedicalStaffsRepository.GetByNikByQuery(ctx, nik)
	if err != nil {
		return []Domain{}, businesses.ErrNotFoundMedicalStaff
	}

	return res, nil
}
