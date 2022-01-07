package repository

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/doctorsRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicalStaffRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/nursesRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/patientRepo"

	"gorm.io/gorm"
)

func NewDoctorRepository(db *gorm.DB) doctorsEntity.Repository {
	return doctorsRepo.NewDoctorsRepository(db)
}

func NewMedicalStaffRepository(db *gorm.DB) medicalStaffEntity.Repository {
	return medicalStaffRepo.NewMedicalStaffRepository(db)
}

func NewPatientRepository(db *gorm.DB) patientEntity.Repository {
	return patientRepo.NewPatientRepository(db)
}

func NewNurseRepository(db *gorm.DB) nursesEntity.Repository {
	return nursesRepo.NewNursesRepository(db)
}
