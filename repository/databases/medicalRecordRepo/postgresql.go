package medicalRecordRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecordRepository struct {
	db *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) medicalRecordEntity.Repository {
	return &MedicalRecordRepository{
		db: db,
	}
}

func (r *MedicalRecordRepository) CreateNewMedicalRecord(ctx context.Context, medicalRecordDomain *medicalRecordEntity.Domain) (*medicalRecordEntity.Domain, error) {
	rec := FromDomain(medicalRecordDomain)
	rec.Uuid, _ = uuid.NewRandom()
	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *MedicalRecordRepository) GetByUuid(ctx context.Context, uuid string) (medicalRecordEntity.Domain, error) {
	rec := MedicalRecord{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return medicalRecordEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *MedicalRecordRepository) UpdateMedicalRecord(ctx context.Context, id string, medicalRecordDomain *medicalRecordEntity.Domain) (*medicalRecordEntity.Domain, error) {
	rec := FromDomain(medicalRecordDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &medicalRecordEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &medicalRecordEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *MedicalRecordRepository) GetByID(ctx context.Context, id uint) (medicalRecordEntity.Domain, error) {
	rec := MedicalRecord{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return medicalRecordEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *MedicalRecordRepository) DeleteMedicalRecordByUuid(ctx context.Context, id string) (string, error) {
	rec := MedicalRecord{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "MedicalRecord Data was Deleted", nil
}

func (r *MedicalRecordRepository) GetMedicalRecords(ctx context.Context) (*[]medicalRecordEntity.Domain, error) {
	var medicalRecord []MedicalRecord
	if err := r.db.Find(&medicalRecord).Error; err != nil {
		return &[]medicalRecordEntity.Domain{}, err
	}
	result := toDomainArray(medicalRecord)
	return &result, nil
}
