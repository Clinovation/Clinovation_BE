package medicineRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/medicineEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) medicineEntity.Repository {
	return &MedicineRepository{
		db: db,
	}
}

func (r *MedicineRepository) CreateNewMedicine(ctx context.Context, medicineDomain *medicineEntity.Domain) (*medicineEntity.Domain, error) {
	rec := FromDomain(medicineDomain)
	rec.Uuid, _ = uuid.NewRandom()
	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *MedicineRepository) GetByID(ctx context.Context, id uint) (medicineEntity.Domain, error) {
	rec := Medicine{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return medicineEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *MedicineRepository) GetByUuid(ctx context.Context, uuid string) (medicineEntity.Domain, error) {
	rec := Medicine{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return medicineEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *MedicineRepository) UpdateMedicine(ctx context.Context, id string, medicineDomain *medicineEntity.Domain) (*medicineEntity.Domain, error) {
	rec := FromDomain(medicineDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &medicineEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &medicineEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *MedicineRepository) DeleteMedicineByUuid(ctx context.Context, id string) (string, error) {
	rec := Medicine{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Medical Staff was Deleted", nil
}

func (r *MedicineRepository) GetMedicine(ctx context.Context) (*[]medicineEntity.Domain, error) {
	var medicine []Medicine
	if err := r.db.Find(&medicine).Error; err != nil {
		return &[]medicineEntity.Domain{}, err
	}
	result := toDomainArray(medicine)
	return &result, nil
}

func (r *MedicineRepository) GetByName(ctx context.Context, name string) ([]medicineEntity.Domain, error) {
	rec := []Medicine{}

	err := r.db.Find(&rec, "name LIKE ?", "%"+name+"%").Error
	if err != nil {
		return nil, err
	}
	result := toDomainArray(rec)

	return result, nil
}
