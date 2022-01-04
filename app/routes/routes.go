package routes

import (
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalStaffController"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	DoctorsController      doctorsController.DoctorController
	MedicalStaffController medicalStaffController.MedicalStaffController
	JWTMiddleware          middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(echo *echo.Echo) {
	//doctor
	doctors := echo.Group("api/v1/doctor")
	doctors.POST("/register", cl.DoctorsController.Registration)
	doctors.POST("/login", cl.DoctorsController.LoginDoctor)

	//doctor with doctor role
	doctor := doctors
	doctor.Use(middleware.JWTWithConfig(cl.JWTMiddleware), DoctorValidation())
	doctors.PUT("/", cl.DoctorsController.UpdateDoctorById)
	doctors.PUT("/uploadAvatar", cl.DoctorsController.UploadAvatar)

	//doctor with doctor or medical staff  role
	doctorAndMedicalStaff := echo.Group("api/v1/doctor")
	doctorAndMedicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), DoctorOrMedicalStaffValidation())
	doctorAndMedicalStaff.GET("/:uuid", cl.DoctorsController.FindDoctorByUuid)
	doctorAndMedicalStaff.GET("/", cl.DoctorsController.GetDoctors)
	doctorAndMedicalStaff.GET("/", cl.DoctorsController.FindDoctorByNameQuery)
	doctorAndMedicalStaff.GET("/", cl.DoctorsController.FindDoctorByNikQuery)
	doctorAndMedicalStaff.DELETE("/", cl.DoctorsController.DeleteDoctorByUuid)

	//medical staff
	medicalStaffs := echo.Group("api/v1/medicalStaff")
	medicalStaffs.POST("/register", cl.MedicalStaffController.Registration)
	medicalStaffs.POST("/login", cl.MedicalStaffController.LoginMedicalStaff)

	//medical staff with medical staff role
	medicalStaff := medicalStaffs
	medicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	medicalStaff.PUT("/", cl.MedicalStaffController.UpdateMedicalStaffById)
	medicalStaff.GET("/:uuid", cl.MedicalStaffController.FindMedicalStaffByUuid)
	medicalStaff.GET("/", cl.MedicalStaffController.GetMedicalStaff)
	medicalStaff.GET("/", cl.MedicalStaffController.FindMedicalStaffByNameQuery)
	medicalStaff.GET("/", cl.MedicalStaffController.FindMedicalStaffByNikQuery)
	medicalStaff.DELETE("/", cl.MedicalStaffController.DeleteMedicalStaffByUuid)
	medicalStaff.PUT("/", cl.MedicalStaffController.UpdateMedicalStaffById)
	medicalStaff.PUT("/uploadAvatar", cl.MedicalStaffController.UploadAvatar)
}

func MedicalStaffValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := auth.GetMedicalStaff(c)

			if user.Role != "medical staff" {
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
			user := auth.GetDoctor(c)

			if user.Role != "doctor" {
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
			user := auth.GetDoctor(c)

			if user.Role != "doctor" && user.Role != "medical staff" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("You are not a Doctor or Medical Staff",
						errors.New("Please Login as Doctor or Medical Stafff"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}
		}
	}
}
