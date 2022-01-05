package main

import (
	"github.com/Clinovation/Clinovation_BE/app/routes"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalStaffController"
	"github.com/Clinovation/Clinovation_BE/controllers/patientController"
	"github.com/Clinovation/Clinovation_BE/helpers"

	ConfigJWT "github.com/Clinovation/Clinovation_BE/app/configs/auth"
	configDB "github.com/Clinovation/Clinovation_BE/app/configs/databases"
	_middleware "github.com/Clinovation/Clinovation_BE/app/middlewares/logger"
	_domainFactory "github.com/Clinovation/Clinovation_BE/repository"

	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	var (
		db  = configDB.SetupDatabaseConnection()
		jwt = ConfigJWT.SetupJwt()
	)
	timeoutDur, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	timeoutContext := time.Duration(timeoutDur) * time.Millisecond

	echoApp := echo.New()

	//middleware
	echoApp.Validator = &helpers.CustomValidator{Validator: validator.New()}
	echoApp.Use(middleware.CORS())
	echoApp.Use(middleware.LoggerWithConfig(_middleware.LoggerConfig()))

	//doctor
	doctorRepo := _domainFactory.NewDoctorRepository(db)
	doctorService := doctorsEntity.NewDoctorsServices(doctorRepo, &jwt, timeoutContext)
	doctorCtrl := doctorsController.NewDoctorController(doctorService, &jwt)

	//medical staff
	medicalStaffRepo := _domainFactory.NewMedicalStaffRepository(db)
	medicalStaffService := medicalStaffEntity.NewMedicalStaffServices(medicalStaffRepo, &jwt, timeoutContext)
	medicalStaffCtrl := medicalStaffController.NewMedicalStaffController(medicalStaffService, &jwt)

	//patient
	patientRepo := _domainFactory.NewPatientRepository(db)
	patientService := patientEntity.NewPatientServices(patientRepo, medicalStaffRepo, &jwt, timeoutContext)
	patientCtrl := patientController.NewPatientsController(patientService, &jwt)

	//routes
	routesInit := routes.ControllerList{
		JWTMiddleware:          jwt.Init(),
		DoctorsController:      *doctorCtrl,
		PatientController:      *patientCtrl,
		MedicalStaffController: *medicalStaffCtrl,
	}
	routesInit.RouteRegister(echoApp)

	port := os.Getenv("PORT")
	log.Fatal(echoApp.Start(":" + port))
}
