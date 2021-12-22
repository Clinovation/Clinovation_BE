package doctorsRepo

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DoctorsRepository struct {
	db *gorm.DB
}

func NewDoctorsRepository(db *gorm.DB) doctorsEntity.Repository {
	return &DoctorsRepository{
		db: db,
	}
}

func (r *DoctorsRepository) CreateNewDoctor(ctx context.Context, doctorDomain *doctorsEntity.Domain) (*doctorsEntity.Domain, error) {
	rec := FromDomain(doctorDomain)
	rec.Uuid, _ = uuid.NewRandom()
	rec.Role = "doctor"

	err := r.db.Create(&rec).Error
	if err != nil {
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *DoctorsRepository) GetByEmail(ctx context.Context, email string) (doctorsEntity.Domain, error) {
	rec := Doctors{}

	err := r.db.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return doctorsEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *DoctorsRepository) GetByUuid(ctx context.Context, uuid string) (doctorsEntity.Domain, error) {
	rec := Doctors{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return doctorsEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *DoctorsRepository) UpdateDoctor(ctx context.Context, id string, doctorDomain *doctorsEntity.Domain) (*doctorsEntity.Domain, error) {
	rec := FromDomain(doctorDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &doctorsEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &doctorsEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *DoctorsRepository) UploadAvatar(ctx context.Context, id string, doctorDomain *doctorsEntity.Domain) (*doctorsEntity.Domain, error) {
	rec := FromDomain(doctorDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &doctorsEntity.Domain{}, err
	}

	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &doctorsEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil
}

func (r *DoctorsRepository) DeleteDoctorByUuid(ctx context.Context, id string) (string, error) {
	rec := Doctors{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "doctor was Deleted", nil
}
