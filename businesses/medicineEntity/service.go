package medicineEntity

import (
	"context"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"time"
)

type MedicineServices struct {
	medicinesRepository Repository
	jwtAuth             *auth.ConfigJWT
	ContextTimeout      time.Duration
}

func NewMedicineServices(repoMedicine Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &MedicineServices{
		medicinesRepository: repoMedicine,
		jwtAuth:             auth,
		ContextTimeout:      timeout,
	}
}

func (qs *MedicineServices) CreateNewMedicine(ctx context.Context, medicineDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	res, err := qs.medicinesRepository.CreateNewMedicine(ctx, medicineDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (qs *MedicineServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	result, err := qs.medicinesRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (qs *MedicineServices) UpdateById(ctx context.Context, medicineDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	result, err := qs.medicinesRepository.UpdateMedicine(ctx, id, medicineDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (qs *MedicineServices) DeleteMedicine(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	res, err := qs.medicinesRepository.DeleteMedicineByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundMedicine
	}
	return res, nil
}

func (qs *MedicineServices) GetMedicines(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	res, err := qs.medicinesRepository.GetMedicine(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}

func (qs *MedicineServices) FindByName(ctx context.Context, name string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	res, err := qs.medicinesRepository.GetByName(ctx, name)
	if err != nil {
		return []Domain{}, businesses.ErrNotFoundMedicine
	}

	return res, nil
}
