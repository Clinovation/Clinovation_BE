package main

import (
	"github.com/Clinovation/Clinovation_BE/app/routes"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController"
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

	//users
	doctorRepo := _domainFactory.NewDoctorRepository(db)
	doctorService := doctorsEntity.NewDoctorsServices(doctorRepo, &jwt, timeoutContext)
	doctorCtrl := doctorsController.NewDoctorController(doctorService, &jwt)

	//routes
	routesInit := routes.ControllerList{
		JWTMiddleware:     jwt.Init(),
		DoctorsController: *doctorCtrl,
	}
	routesInit.RouteRegister(echoApp)

	port := os.Getenv("PORT")
	log.Fatal(echoApp.Start(":" + port))
}
