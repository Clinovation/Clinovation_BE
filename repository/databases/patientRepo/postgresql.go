package patientRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) patientEntity.Repository {
	return &PatientRepository{
		db: db,
	}
}

func (r *PatientRepository) CreateNewPatient(ctx context.Context, patientDomain *patientEntity.Domain) (*patientEntity.Domain, error) {
	rec := FromDomain(patientDomain)
	rec.Uuid, _ = uuid.NewRandom()
	rec.Role = "patient"
	err := r.db.Create(&rec).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := ToDomain(rec)
	return &result, nil
}

func (r *PatientRepository) GetByNik(ctx context.Context, nik string) (patientEntity.Domain, error) {
	rec := Patient{}

	err := r.db.Where("nik = ?", nik).First(&rec).Error
	if err != nil {
		return patientEntity.Domain{}, err
	}

	return ToDomain(&rec), nil
}

func (r *PatientRepository) GetByUuid(ctx context.Context, uuid string) (patientEntity.Domain, error) {
	rec := Patient{}
	err := r.db.Where("uuid = ?", uuid).First(&rec).Error
	if err != nil {
		return patientEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *PatientRepository) UpdatePatient(ctx context.Context, id string, patientDomain *patientEntity.Domain) (*patientEntity.Domain, error) {
	rec := FromDomain(patientDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &patientEntity.Domain{}, err
	}
	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &patientEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil

}

func (r *PatientRepository) GetByID(ctx context.Context, id uint) (patientEntity.Domain, error) {
	rec := Patient{}

	err := r.db.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return patientEntity.Domain{}, err
	}
	return ToDomain(&rec), nil
}

func (r *PatientRepository) UploadAvatar(ctx context.Context, id string, patientDomain *patientEntity.Domain) (*patientEntity.Domain, error) {
	rec := FromDomain(patientDomain)

	if err := r.db.Where("uuid = ?", id).Updates(&rec).Error; err != nil {
		return &patientEntity.Domain{}, err
	}

	if err := r.db.Where("uuid = ?", id).First(&rec).Error; err != nil {
		return &patientEntity.Domain{}, err
	}

	result := ToDomain(rec)

	return &result, nil
}

func (r *PatientRepository) DeletePatientByUuid(ctx context.Context, id string) (string, error) {
	rec := Patient{}

	if err := r.db.Where("uuid = ?", id).Delete(&rec).Error; err != nil {
		return "", err
	}

	return "Patient Data was Deleted", nil
}

func (r *PatientRepository) GetPatients(ctx context.Context, offset, limit int) (*[]patientEntity.Domain, int64, error) {
	var totalData int64
	domain := []patientEntity.Domain{}
	rec := []Patient{}

	r.db.Find(&rec).Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Joins("MedicalStaff").Find(&rec).Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].MedicalStaff = rec[i].MedicalStaff.Name
	}
	return &domain, totalData, nil
}

func (r *PatientRepository) GetByName(ctx context.Context, name string, offset, limit int) ([]patientEntity.Domain, int64, error) {
	var totalData int64
	domain := []patientEntity.Domain{}
	rec := []Patient{}

	r.db.Find(&rec, "name LIKE ?", "%"+name+"%").Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Joins("MedicalStaff").Find(&rec, "patients.name LIKE ?", "%"+name+"%").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].MedicalStaff = rec[i].MedicalStaff.Name
	}
	return domain, totalData, nil
}

func (r *PatientRepository) GetByNikByQuery(ctx context.Context, nik string, offset, limit int) ([]patientEntity.Domain, int64, error) {
	var totalData int64
	domain := []patientEntity.Domain{}
	rec := []Patient{}

	r.db.Find(&rec, "nik LIKE ?", "%"+nik+"%").Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Joins("MedicalStaff").Find(&rec, "patients.nik LIKE ?", "%"+nik+"%").Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].MedicalStaff = rec[i].MedicalStaff.Name
	}
	return domain, totalData, nil
}
