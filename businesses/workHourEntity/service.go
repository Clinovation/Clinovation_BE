package workHourEntity

import (
	"context"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"strings"
	"time"
)

type WorkHoursServices struct {
	WorkHoursRepository Repository
	jwtAuth             *auth.ConfigJWT
	ContextTimeout      time.Duration
}

func NewWorkHoursServices(repoWorkHour Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &WorkHoursServices{
		WorkHoursRepository: repoWorkHour,
		jwtAuth:             auth,
		ContextTimeout:      timeout,
	}
}

func (wds *WorkHoursServices) CreateWorkHour(ctx context.Context, workHourDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	existedWorkHour, err := wds.WorkHoursRepository.GetByHour(ctx, workHourDomain.Hour)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedWorkHour != (Domain{}) {
		return nil, businesses.ErrDuplicateWorkHour
	}

	res, err := wds.WorkHoursRepository.CreateNewWorkHour(ctx, workHourDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (wds *WorkHoursServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	result, err := wds.WorkHoursRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (wds *WorkHoursServices) FindByHour(ctx context.Context, hour string, page int) ([]Domain, int, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	var offset int
	limit := 5
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 5
	}

	res, totalData, err := wds.WorkHoursRepository.GetByHourByQuery(ctx, hour, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, businesses.ErrNotFoundWorkHour
	}

	return res, offset, limit, totalData, nil
}

func (wds *WorkHoursServices) UpdateById(ctx context.Context, workHourDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	result, err := wds.WorkHoursRepository.UpdateWorkHour(ctx, id, workHourDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (wds *WorkHoursServices) DeleteWorkHour(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	res, err := wds.WorkHoursRepository.DeleteWorkHourByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundWorkHour
	}
	return res, nil
}

func (wds *WorkHoursServices) GetWorkHours(ctx context.Context, page int) (*[]Domain, int, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	var offset int
	limit := 5
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 5
	}

	res, totalData, err := wds.WorkHoursRepository.GetWorkHours(ctx, offset, limit)
	if err != nil {
		return &[]Domain{}, -1, -1, -1, businesses.ErrNotFoundWorkHour
	}

	return res, offset, limit, totalData, nil
}
