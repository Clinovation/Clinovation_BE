package doctorsEntity

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

type DoctorsServices struct {
	DoctorsRepository Repository
	jwtAuth           *auth.ConfigJWT
	ContextTimeout    time.Duration
}

func NewDoctorsServices(repoDoctor Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &DoctorsServices{
		DoctorsRepository: repoDoctor,
		jwtAuth:           auth,
		ContextTimeout:    timeout,
	}
}

func (ds *DoctorsServices) Register(ctx context.Context, doctorDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	existedDoctorByNik, err := ds.DoctorsRepository.GetByNik(ctx, doctorDomain.Nik)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedDoctorByNik != (Domain{}) {
		return nil, businesses.ErrDuplicateNik
	}

	existedDoctorByEmail, err := ds.DoctorsRepository.GetByEmail(ctx, doctorDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedDoctorByEmail != (Domain{}) {
		return nil, businesses.ErrDuplicateEmail
	}

	doctorDomain.Password, err = helpers.HashPassword(doctorDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := ds.DoctorsRepository.CreateNewDoctor(ctx, doctorDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (ds *DoctorsServices) Login(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	doctorDomain, err := ds.DoctorsRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", businesses.ErrEmailNotRegistered
	}

	if !helpers.ValidateHash(password, doctorDomain.Password) {
		return "", businesses.ErrPassword
	}

	token := ds.jwtAuth.GenerateToken(doctorDomain.Uuid.String(), doctorDomain.Role)

	return token, nil
}

func (ds *DoctorsServices) Logout(ctx echo.Context) error {
	cookie, err := auth.LogoutCookie(ctx)
	fmt.Println(cookie)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DoctorsServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	result, err := ds.DoctorsRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (ds *DoctorsServices) UpdateById(ctx context.Context, doctorDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	passwordHash, err := helpers.HashPassword(doctorDomain.Password)
	if err != nil {
		panic(err)
	}

	doctorDomain.Password = passwordHash
	result, err := ds.DoctorsRepository.UpdateDoctor(ctx, id, doctorDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (ds *DoctorsServices) UploadAvatar(ctx context.Context, id string, imageLink string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	doctor, err := ds.DoctorsRepository.GetByUuid(ctx, id)
	if err != nil {
		return &Domain{}, err
	}

	doctor.Avatar = imageLink
	updateAvatar, err := ds.DoctorsRepository.UploadAvatar(ctx, id, &doctor)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}

func (ds *DoctorsServices) DeleteDoctor(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	res, err := ds.DoctorsRepository.DeleteDoctorByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundDoctor
	}
	return res, nil
}

func (ds *DoctorsServices) GetDoctors(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	res, err := ds.DoctorsRepository.GetDoctors(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}
