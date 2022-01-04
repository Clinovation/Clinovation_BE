package medicalStaffRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalStaffRepository struct {
	db *gorm.DB
}

func NewMedicalStaffRepository(db *gorm.DB) medicalStaffEntity.Repository {
	return &MedicalStaffRepository{
		db: db,
	}
}

func (r *MedicalStaffRepository) CreateNewMedicalStaff(ctx context.Context, medicalStaffDomain *medicalStaffEntity.Domain) (*medicalStaffEntity.Domain, error) {
	rec := FromDomain(medicalStaffDomain)
	rec.Uuid, _ = uuid.NewRandom()
	rec.Role = "medical staff"
	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *MedicalStaffRepository) GetByNik(ctx context.Context, nik string) (medicalStaffEntity.Domain, error) {
	rec := MedicalStaff{}

	err := r.db.Where("nik = ?", nik).First(&rec).Error
	if err != nil {
		return medicalStaffEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *MedicalStaffRepository) GetByEmail(ctx context.Context, email string) (medicalStaffEntity.Domain, error) {
	rec := MedicalStaff{}

	err := r.db.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return medicalStaffEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *MedicalStaffRepository) GetByUuid(ctx context.Context, uuid string) (medicalStaffEntity.Domain, error) {
	rec := MedicalStaff{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return medicalStaffEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *MedicalStaffRepository) UpdateMedicalStaff(ctx context.Context, id string, medicalStaffDomain *medicalStaffEntity.Domain) (*medicalStaffEntity.Domain, error) {
	rec := FromDomain(medicalStaffDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &medicalStaffEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &medicalStaffEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *MedicalStaffRepository) UploadAvatar(ctx context.Context, id string, medicalStaffDomain *medicalStaffEntity.Domain) (*medicalStaffEntity.Domain, error) {
	rec := FromDomain(medicalStaffDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &medicalStaffEntity.Domain{}, err
	}

	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &medicalStaffEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil
}

func (r *MedicalStaffRepository) DeleteMedicalStaffByUuid(ctx context.Context, id string) (string, error) {
	rec := MedicalStaff{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Medical Staff was Deleted", nil
}

func (r *MedicalStaffRepository) GetMedicalStaff(ctx context.Context) (*[]medicalStaffEntity.Domain, error) {
	var medicalStaff []MedicalStaff
	if err := r.db.Find(&medicalStaff).Error; err != nil {
		return &[]medicalStaffEntity.Domain{}, err
	}
	result := toDomainArray(medicalStaff)
	return &result, nil
}

func (r *MedicalStaffRepository) GetByName(ctx context.Context, name string) ([]medicalStaffEntity.Domain, error) {
	rec := []MedicalStaff{}

	err := r.db.Find(&rec, "name LIKE ?", "%"+name+"%").Error
	if err != nil {
		return nil, err
	}
	result := toDomainArray(rec)

	return result, nil
}

func (r *MedicalStaffRepository) GetByNikByQuery(ctx context.Context, nik string) ([]medicalStaffEntity.Domain, error) {
	rec := []MedicalStaff{}

	err := r.db.Find(&rec, "nik LIKE ?", "%"+nik+"%").Error
	if err != nil {
		return nil, err
	}
	result := toDomainArray(rec)

	return result, nil
}
