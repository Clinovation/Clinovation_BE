package routes

import (
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	DoctorsController doctorsController.DoctorController
	JWTMiddleware     middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(echo *echo.Echo) {
	//doctor
	doctors := echo.Group("api/v1/doctor")
	doctors.POST("/register", cl.DoctorsController.Registration)
	doctors.POST("/login", cl.DoctorsController.LoginDoctor)
	doctors.POST("/logout", cl.DoctorsController.LogoutDoctor)
	doctors.GET("/:uuid", cl.DoctorsController.FindDoctorByUuid)

	//doctor with doctor role
	doctor := doctors
	doctor.Use(middleware.JWTWithConfig(cl.JWTMiddleware), DoctorValidation())
	doctors.PUT("/", cl.DoctorsController.UpdateDoctorById)
	doctors.GET("/", cl.DoctorsController.GetDoctors)
	doctors.DELETE("/", cl.DoctorsController.DeleteDoctorByUuid)
	doctors.PUT("/uploadAvatar", cl.DoctorsController.UploadAvatar)
}

func MedicalStaffValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			medicalStaff := auth.GetMedicalStaff(c)

			if medicalStaff.Role != "medical staff" {
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
			doctor := auth.GetDoctor(c)

			if doctor.Role != "doctor" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("You are not a Doctor",
						errors.New("Please Login as Doctor"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}

		}
	}
}

func DoctorOrMedicalStaffValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			doctor := auth.GetDoctor(c)

			if doctor.Role != "doctor" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("You are not a Doctor",
						errors.New("Please Login as Doctor"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}

		}
	}
}
