package repository

import (
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicineEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/queueEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/scheduleEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/workHourEntity"
	"github.com/Clinovation/Clinovation_BE/repository/databases/doctorsRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicalStaffRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicineRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/nursesRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/patientRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/queueRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/scheduleRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/workDayRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/workHourRepo"

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

func NewWorkDayRepository(db *gorm.DB) workDayEntity.Repository {
	return workDayRepo.NewWorkDaysRepository(db)
}

func NewWorkHourRepository(db *gorm.DB) workHourEntity.Repository {
	return workHourRepo.NewWorkHoursRepository(db)
}

func NewScheduleRepository(db *gorm.DB) scheduleEntity.Repository {
	return scheduleRepo.NewScheduleRepository(db)
}

func NewQueueRepository(db *gorm.DB) queueEntity.Repository {
	return queueRepo.NewQueueRepository(db)
}

func NewMedicineRepository(db *gorm.DB) medicineEntity.Repository {
	return medicineRepo.NewMedicineRepository(db)
}
