package routes

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	DoctorsController    doctorsController.DoctorController
	JWTMiddleware      middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(echo *echo.Echo) {
	//doctor
	doctors := echo.Group("api/v1/doctor")
	doctors.POST("/register", cl.DoctorsController.Registration)
	doctors.POST("/login", cl.DoctorsController.LoginDoctor)
	doctors.POST("/logout", cl.DoctorsController.LogoutDoctor)
	doctors.GET("/:uuid", cl.DoctorsController.FindDoctorByUuid)

	//doctor with admin role
	doctors.PUT("/", cl.DoctorsController.UpdateDoctorById, middleware.JWTWithConfig(cl.JWTMiddleware))
	doctors.DELETE("/", cl.DoctorsController.DeleteDoctorByUuid, middleware.JWTWithConfig(cl.JWTMiddleware))
	doctors.POST("/uploadAvatar", cl.DoctorsController.UploadAvatar, middleware.JWTWithConfig(cl.JWTMiddleware))
}

func MedicalStaffValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := auth.GetUser(c)

			if user.Role != "admin" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("you are not a Medical Staff",
						errors.New("Please Login as Medical Staff"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}

		}
	}
}

func DoctorValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := auth.GetUser(c)

			if user.Role != "admin" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("you are not a Doctor",
						errors.New("Please Login as Doctor"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}

		}
	}
}

func NurseValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := auth.GetUser(c)

			if user.Role != "admin" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("you are not a Nurse",
						errors.New("Please Login as Nurse"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}

		}
	}
}
