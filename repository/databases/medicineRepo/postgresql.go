package medicineRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/medicineEntity"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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

func (r *MedicineRepository) GetMedicine(ctx context.Context, offset, limit int) (*[]medicineEntity.Domain, int64, error) {
	var totalData int64
	domain := []medicineEntity.Domain{}
	rec := []Medicine{}

	r.db.Find(&rec).Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Find(&rec).Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)

	return &domain, totalData, nil
}

func (r *MedicineRepository) GetByNameByQuery(ctx context.Context, name string, offset, limit int) ([]medicineEntity.Domain, int64, error) {
	var totalData int64
	domain := []medicineEntity.Domain{}
	rec := []Medicine{}

	r.db.Find(&rec, "name LIKE ?", "%"+name+"%").Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Find(&rec, "name LIKE ?", "%"+name+"%").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)

	return domain, totalData, nil
}

func (r *MedicineRepository) GetByName(ctx context.Context, name string) (medicineEntity.Domain, error) {
	rec := Medicine{}

	err := r.db.Where("name = ?", name).First(&rec).Error
	if err != nil {
		return medicineEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}
