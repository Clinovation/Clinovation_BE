package nursesEntity

import (
	"context"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"strings"
	"time"
)

type NursesServices struct {
	NursesRepository Repository
	jwtAuth          *auth.ConfigJWT
	ContextTimeout   time.Duration
}

func NewNursesServices(repoNurse Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &NursesServices{
		NursesRepository: repoNurse,
		jwtAuth:          auth,
		ContextTimeout:   timeout,
	}
}

func (ns *NursesServices) Register(ctx context.Context, nurserDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	existedNurseByNik, err := ns.NursesRepository.GetByNik(ctx, nurserDomain.Nik)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedNurseByNik != (Domain{}) {
		return nil, businesses.ErrDuplicateNik
	}

	existedNurseByEmail, err := ns.NursesRepository.GetByEmail(ctx, nurserDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedNurseByEmail != (Domain{}) {
		return nil, businesses.ErrDuplicateEmail
	}

	nurserDomain.Password, err = helpers.HashPassword(nurserDomain.Password)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	res, err := ns.NursesRepository.CreateNewNurse(ctx, nurserDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (ns *NursesServices) Login(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	nurserDomain, err := ns.NursesRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", businesses.ErrEmailNotRegistered
	}

	if !helpers.ValidateHash(password, nurserDomain.Password) {
		return "", businesses.ErrPassword
	}

	token := ns.jwtAuth.GenerateToken(nurserDomain.Uuid.String(), nurserDomain.Role)

	return token, nil
}

func (ns *NursesServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	result, err := ns.NursesRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (ds *NursesServices) ChangePassword(ctx context.Context, nurserDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	passwordHash, err := helpers.HashPassword(nurserDomain.Password)
	if err != nil {
		panic(err)
	}

	nurserDomain.Password = passwordHash
	result, err := ds.NursesRepository.UpdateNurse(ctx, id, nurserDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (ds *NursesServices) ForgetPassword(ctx context.Context, nurserDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ds.ContextTimeout)
	defer cancel()

	result, err := ds.NursesRepository.ForgetPassword(ctx, nurserDomain.Nik, nurserDomain.Email)
	if err != nil {
		return &Domain{}, err
	}
	return &result, nil
}

func (ns *NursesServices) UpdateById(ctx context.Context, nurserDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	passwordHash, err := helpers.HashPassword(nurserDomain.Password)
	if err != nil {
		panic(err)
	}

	nurserDomain.Password = passwordHash
	result, err := ns.NursesRepository.UpdateNurse(ctx, id, nurserDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (ns *NursesServices) UploadAvatar(ctx context.Context, id string, imageLink string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	nurse, err := ns.NursesRepository.GetByUuid(ctx, id)
	if err != nil {
		return &Domain{}, err
	}

	nurse.Avatar = imageLink
	updateAvatar, err := ns.NursesRepository.UploadAvatar(ctx, id, &nurse)
	if err != nil {
		return &Domain{}, err
	}
	return updateAvatar, nil
}

func (ns *NursesServices) DeleteNurse(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	res, err := ns.NursesRepository.DeleteNurseByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundNurse
	}
	return res, nil
}

func (ns *NursesServices) GetNurses(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	res, err := ns.NursesRepository.GetNurses(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}

func (ns *NursesServices) FindByName(ctx context.Context, name string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	res, err := ns.NursesRepository.GetByName(ctx, name)
	if err != nil {
		return []Domain{}, businesses.ErrNotFoundNurse
	}

	return res, nil
}

func (ns *NursesServices) FindByNik(ctx context.Context, nik string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, ns.ContextTimeout)
	defer cancel()

	res, err := ns.NursesRepository.GetByNikByQuery(ctx, nik)
	if err != nil {
		return []Domain{}, businesses.ErrNotFoundNurse
	}

	return res, nil
}
