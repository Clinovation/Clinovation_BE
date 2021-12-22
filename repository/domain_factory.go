package repository

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/doctorsRepo"

	"gorm.io/gorm"
)

func NewDoctorRepository(db *gorm.DB) doctorsEntity.Repository {
	return doctorsRepo.NewDoctorsRepository(db)
}

