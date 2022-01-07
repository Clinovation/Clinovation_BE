package routes

import (
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalStaffController"
	"github.com/Clinovation/Clinovation_BE/controllers/nursesController"
	"github.com/Clinovation/Clinovation_BE/controllers/patientController"
	"github.com/Clinovation/Clinovation_BE/controllers/workDayController"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	DoctorsController      doctorsController.DoctorController
	NurseController        nursesController.NurseController
	MedicalStaffController medicalStaffController.MedicalStaffController
	PatientController      patientController.PatientsController
	WorkDayController      workDayController.WorkDayController
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
	doctor.PUT("/", cl.DoctorsController.UpdateDoctorById)
	doctor.PUT("/uploadAvatar", cl.DoctorsController.UploadAvatar)

	//doctor with doctor or medical staff  role
	doctorAndMedicalStaff := echo.Group("api/v1/doctor")
	doctorAndMedicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), DoctorOrMedicalStaffValidation())
	doctorAndMedicalStaff.GET("/:uuid", cl.DoctorsController.FindDoctorByUuid)
	doctorAndMedicalStaff.GET("/", cl.DoctorsController.GetDoctors)
	doctorAndMedicalStaff.GET("/", cl.DoctorsController.FindDoctorByNameQuery)
	doctorAndMedicalStaff.GET("/", cl.DoctorsController.FindDoctorByNikQuery)
	doctorAndMedicalStaff.DELETE("/", cl.DoctorsController.DeleteDoctorByUuid)

	//nurse
	nurses := echo.Group("api/v1/nurse")
	nurses.POST("/register", cl.NurseController.Registration)
	nurses.POST("/login", cl.NurseController.LoginNurse)

	//nurse with nurse role
	nurse := nurses
	nurse.Use(middleware.JWTWithConfig(cl.JWTMiddleware), NurseValidation())
	nurse.PUT("/", cl.NurseController.UpdateNurseById)
	nurse.PUT("/uploadAvatar", cl.NurseController.UploadAvatar)

	//nurse with nurse or medical staff  role
	nurseAndMedicalStaff := echo.Group("api/v1/nurse")
	nurseAndMedicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), NurseOrMedicalStaffValidation())
	nurseAndMedicalStaff.GET("/:uuid", cl.NurseController.FindNurseByUuid)
	nurseAndMedicalStaff.GET("/", cl.NurseController.GetNurses)
	nurseAndMedicalStaff.GET("/", cl.NurseController.FindNurseByNameQuery)
	nurseAndMedicalStaff.GET("/", cl.NurseController.FindNurseByNikQuery)
	nurseAndMedicalStaff.DELETE("/", cl.NurseController.DeleteNurseByUuid)

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
	medicalStaff.PUT("/uploadAvatar", cl.MedicalStaffController.UploadAvatar)

	//patient with medical staff role
	patientMedicalStaff := echo.Group("api/v1/patient")
	patientMedicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	patientMedicalStaff.POST("/register", cl.PatientController.Registration)
	patientMedicalStaff.PUT("/:uuid", cl.PatientController.UpdatePatientById)
	patientMedicalStaff.DELETE("/:uuid", cl.PatientController.DeletePatientByUuid)
	patientMedicalStaff.PUT("/uploadAvatar/:uuid", cl.PatientController.UploadAvatar)

	//patient with doctor,medical staff and nurse role
	patientAllRole := echo.Group("api/v1/patient")
	patientAllRole.Use(middleware.JWTWithConfig(cl.JWTMiddleware), AllRole())
	patientAllRole.GET("/:uuid", cl.PatientController.FindPatientByUuid)
	patientAllRole.GET("/", cl.PatientController.GetPatients)
	patientAllRole.GET("/", cl.PatientController.FindPatientByNameQuery)
	patientAllRole.GET("/", cl.PatientController.FindPatientByNikQuery)

	//work day with medical staff role
	workDays := echo.Group("api/v1/workDay")
	workDays.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	workDays.POST("/", cl.WorkDayController.CreateNewWorkDay)
	workDays.PUT("/:uuid", cl.WorkDayController.UpdateWorkDayById)
	workDays.GET("/:uuid", cl.WorkDayController.FindWorkDayByUuid)
	workDays.GET("/", cl.WorkDayController.GetWorkDays)
	workDays.GET("/:day", cl.WorkDayController.FindWorkDayByDay)
	workDays.DELETE("/:uuid", cl.WorkDayController.DeleteWorkDayByUuid)

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
			user := auth.GetUser(c)

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

func NurseValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := auth.GetNurse(c)

			if user.Role != "nurse" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("You are not a Nurse",
						errors.New("Please Login as Nurse"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}

		}
	}
}

func NurseOrMedicalStaffValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := auth.GetUser(c)

			if user.Role != "nurse" && user.Role != "medical staff" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("You are not a Nurse or Medical Staff",
						errors.New("Please Login as Nurse or Medical Stafff"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}
		}
	}
}

func AllRole() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := auth.GetUser(c)

			if user.Role != "nurse" && user.Role != "medical staff" && user.Role != "doctor" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("You are not a Nurse or Medical Staff or doctor",
						errors.New("Please Login as Nurse or Medical Stafff or doctor"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}
		}
	}
}
