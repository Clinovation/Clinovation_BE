package config

import (
	"fmt"
	"github.com/Clinovation/Clinovation_BE/repository/databases/nursesRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/workDayRepo"
	"os"

	"github.com/Clinovation/Clinovation_BE/repository/databases/doctorsRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/medicalStaffRepo"
	"github.com/Clinovation/Clinovation_BE/repository/databases/patientRepo"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabaseConnection() *gorm.DB {
	_ = godotenv.Load()

	//var dbName string
	//if os.Getenv("ENV") == "TESTING"{
	//	dbName = os.Getenv("DB_NAME_TESTING")
	//} else {
	//	dbName = os.Getenv("DB_NAME")
	//}
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// https://github.com/go-gorm/postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		//PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err.Error())
	}
	dbMigrate(db)

	return db
}

func dbMigrate(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(
		&doctorsRepo.Doctors{},
		&medicalStaffRepo.MedicalStaff{},
		&patientRepo.Patient{},
		&nursesRepo.Nurses{},
		&workDayRepo.WorkDays{},
	)
}
