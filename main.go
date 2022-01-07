package main

import (
	"github.com/Clinovation/Clinovation_BE/app/routes"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
	"github.com/Clinovation/Clinovation_BE/businesses/workHourEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalStaffController"
	"github.com/Clinovation/Clinovation_BE/controllers/nursesController"
	"github.com/Clinovation/Clinovation_BE/controllers/patientController"
	"github.com/Clinovation/Clinovation_BE/controllers/workDayController"
	"github.com/Clinovation/Clinovation_BE/controllers/workHourController"
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

	//nurse
	nurseRepo := _domainFactory.NewNurseRepository(db)
	nurseService := nursesEntity.NewNursesServices(nurseRepo, &jwt, timeoutContext)
	nurseCtrl := nursesController.NewNursesController(nurseService, &jwt)

	//medical staff
	medicalStaffRepo := _domainFactory.NewMedicalStaffRepository(db)
	medicalStaffService := medicalStaffEntity.NewMedicalStaffServices(medicalStaffRepo, &jwt, timeoutContext)
	medicalStaffCtrl := medicalStaffController.NewMedicalStaffController(medicalStaffService, &jwt)

	//patient
	patientRepo := _domainFactory.NewPatientRepository(db)
	patientService := patientEntity.NewPatientServices(patientRepo, medicalStaffRepo, &jwt, timeoutContext)
	patientCtrl := patientController.NewPatientsController(patientService, &jwt)

	//work Day
	workDayRepo := _domainFactory.NewWorkDayRepository(db)
	workDayService := workDayEntity.NewWorkDaysServices(workDayRepo, &jwt, timeoutContext)
	workDayCtrl := workDayController.NewWorkDayController(workDayService, &jwt)

	//work Hour
	workHourRepo := _domainFactory.NewWorkHourRepository(db)
	workHourService := workHourEntity.NewWorkHoursServices(workHourRepo, &jwt, timeoutContext)
	workHourCtrl := workHourController.NewWorkHourController(workHourService, &jwt)

	//routes
	routesInit := routes.ControllerList{
		JWTMiddleware:          jwt.Init(),
		DoctorsController:      *doctorCtrl,
		NurseController:        *nurseCtrl,
		PatientController:      *patientCtrl,
		WorkDayController:      *workDayCtrl,
		WorkHourController:     *workHourCtrl,
		MedicalStaffController: *medicalStaffCtrl,
	}
	routesInit.RouteRegister(echoApp)

	port := os.Getenv("PORT")
	log.Fatal(echoApp.Start(":" + port))
}
