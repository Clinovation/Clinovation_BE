package medicalRecordRepo

import (
	"context"
	"fmt"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
	domain := medicalRecordEntity.Domain{}
	rec := MedicalRecord{}
	err := r.db.Joins("MedicalStaff").Joins("Patient").Find(&rec, "medical_records.uuid = ?", uuid).Error
	if err != nil {
		return medicalRecordEntity.Domain{}, err
	}

	copier.Copy(&domain, &rec)

	domain.MedicalStaff = rec.MedicalStaff.Name
	domain.PatientName = rec.Patient.Name
	domain.PatientAddress = rec.Patient.Address
	domain.PatientDob = rec.Patient.Dob
	domain.PatientHeight = rec.Patient.Height
	domain.PatientWeight = rec.Patient.Weight
	domain.PatientNik = rec.Patient.Nik
	domain.PatientSex = rec.Patient.Sex
	domain.PatientStatusMartial = rec.Patient.StatusMartial

	//return ToDomain(&rec), nil
	return domain, nil
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

func (r *MedicalRecordRepository) GetMedicalRecordsQueue(ctx context.Context, userID uint, offset, limit int) (*[]medicalRecordEntity.Domain, int64, error) {
	var totalData int64
	domain := []medicalRecordEntity.Domain{}
	rec := []MedicalRecord{}

	r.db.Find(&rec, "user_id = ?", userID).Count(&totalData)
	err := r.db.Limit(limit).Offset(offset).Joins("MedicalStaff").Joins("Patient").Find(&rec, "user_id = ?", userID).Error
	if err != nil {
		return nil, 0, err
	}

	copier.Copy(&domain, &rec)
	for i := 0; i < len(rec); i++ {
		domain[i].MedicalStaff = rec[i].MedicalStaff.Name
		domain[i].PatientName = rec[i].Patient.Name
		domain[i].PatientAddress = rec[i].Patient.Address
		domain[i].PatientDob = rec[i].Patient.Dob
		domain[i].PatientHeight = rec[i].Patient.Height
		domain[i].PatientWeight = rec[i].Patient.Weight
		domain[i].PatientNik = rec[i].Patient.Nik
		domain[i].PatientSex = rec[i].Patient.Sex
		domain[i].PatientStatusMartial = rec[i].Patient.StatusMartial
	}
	return &domain, totalData, nil
}
