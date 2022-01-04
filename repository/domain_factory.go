package repository

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/doctorsRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicalStaffRepo"

	"gorm.io/gorm"
)

func NewDoctorRepository(db *gorm.DB) doctorsEntity.Repository {
	return doctorsRepo.NewDoctorsRepository(db)
}

func NewMedicalStaffRepository(db *gorm.DB) medicalStaffEntity.Repository {
	return medicalStaffRepo.NewMedicalStaffRepository(db)
}


