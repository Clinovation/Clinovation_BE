package workDayEntity

import (
	"context"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"strings"
	"time"
)

type WorkDaysServices struct {
	WorkDaysRepository Repository
	jwtAuth            *auth.ConfigJWT
	ContextTimeout     time.Duration
}

func NewWorkDaysServices(repoWorkDay Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &WorkDaysServices{
		WorkDaysRepository: repoWorkDay,
		jwtAuth:            auth,
		ContextTimeout:     timeout,
	}
}

func (wds *WorkDaysServices) CreateWorkDay(ctx context.Context, workDayDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	existedWorkDay, err := wds.WorkDaysRepository.GetByDay(ctx, workDayDomain.Day)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	if existedWorkDay != (Domain{}) {
		return nil, businesses.ErrDuplicateWorkDay
	}

	res, err := wds.WorkDaysRepository.CreateNewWorkDay(ctx, workDayDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (wds *WorkDaysServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	result, err := wds.WorkDaysRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (wds *WorkDaysServices) FindByDay(ctx context.Context, day string, page int) ([]Domain, int, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	var offset int
	limit := 5
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 5
	}

	res, totalData, err := wds.WorkDaysRepository.GetByDayByQuery(ctx, day, offset, limit)
	if err != nil {
		return []Domain{}, -1, -1, -1, businesses.ErrNotFoundWorkDay
	}

	return res, offset, limit, totalData, nil
}

func (wds *WorkDaysServices) UpdateById(ctx context.Context, workDayDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	result, err := wds.WorkDaysRepository.UpdateWorkDay(ctx, id, workDayDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (wds *WorkDaysServices) DeleteWorkDay(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	res, err := wds.WorkDaysRepository.DeleteWorkDayByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundWorkDay
	}
	return res, nil
}

func (wds *WorkDaysServices) GetWorkDaysPagination(ctx context.Context, page int) (*[]Domain, int, int, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	var offset int
	limit := 5
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * 5
	}

	res, totalData, err := wds.WorkDaysRepository.GetWorkDaysPagination(ctx, offset, limit)
	if err != nil {
		return &[]Domain{}, -1, -1, -1, businesses.ErrNotFoundWorkDay
	}

	return res, offset, limit, totalData, nil
}

func (wds *WorkDaysServices) GetWorkDays(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, wds.ContextTimeout)
	defer cancel()

	res, err := wds.WorkDaysRepository.GetWorkDays(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}
