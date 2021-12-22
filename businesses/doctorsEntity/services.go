package doctorsEntity

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)


type DoctorsServices struct {
	DoctorsRepository Repository
	jwtAuth        *auth.ConfigJWT
	ContextTimeout time.Duration
}

func NewDoctorsServices(repoDoctor Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &DoctorsServices{
		DoctorsRepository: repoDoctor,
		jwtAuth:        auth,
		ContextTimeout: timeout,
	}
}

func (us *DoctorsServices) Register(ctx context.Context, doctorDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	existedDoctor, err := us.DoctorsRepository.GetByEmail(ctx, doctorDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedDoctor != (Domain{}) {
		return nil, businesses.ErrDuplicateEmail
	}

	doctorDomain.Password, err = helpers.HashPassword(doctorDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := us.DoctorsRepository.CreateNewDoctor(ctx, doctorDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (us *DoctorsServices) Login(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	doctorDomain, err := us.DoctorsRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", businesses.ErrEmailNotRegistered
	}

	if !helpers.ValidateHash(password, doctorDomain.Password) {
		return "", businesses.ErrPassword
	}

	token := us.jwtAuth.GenerateToken(doctorDomain.Uuid.String(), doctorDomain.Role)

	return token, nil
}

func (us *DoctorsServices) Logout(ctx echo.Context) error {
	cookie, err := auth.LogoutCookie(ctx)
	fmt.Println(cookie)
	if err != nil {
		return err
	}

	return nil
}

func (us *DoctorsServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	result, err := us.DoctorsRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (us *DoctorsServices) UpdateById(ctx context.Context, doctorDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	passwordHash, err := helpers.HashPassword(doctorDomain.Password)
	if err != nil {
		panic(err)
	}

	doctorDomain.Password = passwordHash
	result, err := us.DoctorsRepository.UpdateDoctor(ctx, id, doctorDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (us *DoctorsServices) UploadAvatar(ctx context.Context, id string, fileLocation string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	doctor, err := us.DoctorsRepository.GetByUuid(ctx, id)
	if err != nil {
		return &Domain{}, err
	}

	doctor.Avatar = fileLocation
	updateAvatar, err := us.DoctorsRepository.UploadAvatar(ctx, id, &doctor)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}

func (us *DoctorsServices) DeleteDoctor(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	res, err := us.DoctorsRepository.DeleteDoctorByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundDoctor
	}
	return res, nil
}
