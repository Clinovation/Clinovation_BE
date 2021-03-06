package routes

import (
	"errors"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalRecordController"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalStaffController"
	"github.com/Clinovation/Clinovation_BE/controllers/medicineController"
	"github.com/Clinovation/Clinovation_BE/controllers/nursesController"
	"github.com/Clinovation/Clinovation_BE/controllers/patientController"
	"github.com/Clinovation/Clinovation_BE/controllers/recipeController"
	"github.com/Clinovation/Clinovation_BE/controllers/workDayController"
	"github.com/Clinovation/Clinovation_BE/controllers/workHourController"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	DoctorsController       doctorsController.DoctorController
	NurseController         nursesController.NurseController
	MedicalStaffController  medicalStaffController.MedicalStaffController
	PatientController       patientController.PatientsController
	WorkDayController       workDayController.WorkDayController
	WorkHourController      workHourController.WorkHourController
	MedicineController      medicineController.MedicineController
	RecipeController        recipeController.RecipeController
	MedicalRecordController medicalRecordController.MedicalRecordsController
	JWTMiddleware           middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(echo *echo.Echo) {
	//doctor
	doctors := echo.Group("api/v1/doctor")
	doctors.POST("/register", cl.DoctorsController.Registration)
	doctors.POST("/login", cl.DoctorsController.LoginDoctor)
	doctors.GET("/forgetPassword", cl.DoctorsController.ForgetPassword)
	doctors.PUT("/changePassword", cl.DoctorsController.ChangePassword)

	//doctor with medical staff role
	doctorMedicalStaff := echo.Group("api/v1/doctor")
	doctorMedicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	doctorMedicalStaff.PUT("/accept/:uuid", cl.DoctorsController.AcceptDoctor)
	doctorMedicalStaff.GET("/waitingList", cl.DoctorsController.GetWaitingList)
	doctorMedicalStaff.DELETE("/:uuid", cl.DoctorsController.DeleteDoctorByMedicalStaff)

	//doctor with doctor role
	doctor := doctors
	doctor.Use(middleware.JWTWithConfig(cl.JWTMiddleware), DoctorValidation())
	doctor.PUT("/", cl.DoctorsController.UpdateDoctorById)
	doctor.GET("/jwt", cl.DoctorsController.FindByJwt)
	doctor.PUT("/uploadAvatar", cl.DoctorsController.UploadAvatar)

	//doctor with doctor or medical staff  role
	doctorAndMedicalStaff := echo.Group("api/v1/doctor")
	doctorAndMedicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), DoctorOrMedicalStaffValidation())
	doctorAndMedicalStaff.GET("/:uuid", cl.DoctorsController.FindDoctorByUuid)
	doctorAndMedicalStaff.GET("/", cl.DoctorsController.GetDoctors)
	doctorAndMedicalStaff.GET("/queryName", cl.DoctorsController.FindDoctorByNameQuery)
	doctorAndMedicalStaff.GET("/queryNik", cl.DoctorsController.FindDoctorByNikQuery)
	doctorAndMedicalStaff.GET("/queryDay", cl.DoctorsController.FindDoctorByDayQuery)
	doctorAndMedicalStaff.DELETE("/", cl.DoctorsController.DeleteDoctorByUuid)

	//nurse
	nurses := echo.Group("api/v1/nurse")
	nurses.POST("/register", cl.NurseController.Registration)
	nurses.POST("/login", cl.NurseController.LoginNurse)
	nurses.GET("/forgetPassword", cl.NurseController.ForgetPassword)
	nurses.PUT("/changePassword", cl.NurseController.ChangePassword)

	//nurse with medical staff role
	nurseMedicalStaff := echo.Group("api/v1/nurse")
	nurseMedicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	nurseMedicalStaff.PUT("/accept/:uuid", cl.NurseController.AcceptNurse)
	nurseMedicalStaff.GET("/waitingList", cl.NurseController.GetWaitingList)
	nurseMedicalStaff.GET("/queryDay", cl.NurseController.FindNurseByDayQuery)
	nurseMedicalStaff.DELETE("/:uuid", cl.NurseController.DeleteNurseByMedicalStaff)

	//nurse with nurse role
	nurse := nurses
	nurse.Use(middleware.JWTWithConfig(cl.JWTMiddleware), NurseValidation())
	nurse.GET("/jwt", cl.NurseController.FindByJwt)
	nurse.PUT("/", cl.NurseController.UpdateNurseById)
	nurse.PUT("/uploadAvatar", cl.NurseController.UploadAvatar)

	//nurse with nurse or medical staff  role
	nurseAndMedicalStaff := echo.Group("api/v1/nurse")
	nurseAndMedicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), NurseOrMedicalStaffValidation())
	nurseAndMedicalStaff.GET("/:uuid", cl.NurseController.FindNurseByUuid)
	nurseAndMedicalStaff.GET("/", cl.NurseController.GetNurses)
	nurseAndMedicalStaff.GET("/queryName", cl.NurseController.FindNurseByNameQuery)
	nurseAndMedicalStaff.GET("/queryNik", cl.NurseController.FindNurseByNikQuery)
	nurseAndMedicalStaff.DELETE("/", cl.NurseController.DeleteNurseByUuid)

	//medical staff
	medicalStaffs := echo.Group("api/v1/medicalStaff")
	medicalStaffs.POST("/register", cl.MedicalStaffController.Registration)
	medicalStaffs.POST("/login", cl.MedicalStaffController.LoginMedicalStaff)
	medicalStaffs.GET("/forgetPassword", cl.MedicalStaffController.ForgetPassword)
	medicalStaffs.PUT("/changePassword", cl.MedicalStaffController.ChangePassword)

	//medical staff with medical staff role
	medicalStaff := medicalStaffs
	medicalStaff.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	medicalStaff.GET("/jwt", cl.MedicalStaffController.FindByJwt)
	medicalStaff.PUT("/", cl.MedicalStaffController.UpdateMedicalStaffById)
	medicalStaff.GET("/:uuid", cl.MedicalStaffController.FindMedicalStaffByUuid)
	medicalStaff.GET("/", cl.MedicalStaffController.GetMedicalStaff)
	medicalStaff.GET("/queryName", cl.MedicalStaffController.FindMedicalStaffByNameQuery)
	medicalStaff.GET("/queryNik", cl.MedicalStaffController.FindMedicalStaffByNikQuery)
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
	patientAllRole.GET("/queryName", cl.PatientController.FindPatientByNameQuery)
	patientAllRole.GET("/queryNik", cl.PatientController.FindPatientByNikQuery)

	//work day with medical staff role
	workDays := echo.Group("api/v1/workDay")
	workDays.GET("/", cl.WorkDayController.GetWorkDays)
	workDays.GET("/:uuid", cl.WorkDayController.FindWorkDayByUuid)
	workDays.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	workDays.POST("/", cl.WorkDayController.CreateNewWorkDay)
	workDays.PUT("/:uuid", cl.WorkDayController.UpdateWorkDayById)
	workDays.GET("/pagination", cl.WorkDayController.GetWorkDaysPagination)
	workDays.GET("/queryDay", cl.WorkDayController.FindWorkDayByDay)
	workDays.DELETE("/:uuid", cl.WorkDayController.DeleteWorkDayByUuid)

	//work Hour with medical staff role
	workHours := echo.Group("api/v1/workHour")
	workHours.GET("/", cl.WorkHourController.GetWorkHours)
	workHours.GET("/:uuid", cl.WorkHourController.FindWorkHourByUuid)
	workHours.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	workHours.POST("/", cl.WorkHourController.CreateNewWorkHour)
	workHours.PUT("/:uuid", cl.WorkHourController.UpdateWorkHourById)
	workHours.GET("/pagination", cl.WorkHourController.GetWorkHoursPagination)
	workHours.GET("/queryHour", cl.WorkHourController.FindWorkHourByHour)
	workHours.DELETE("/:uuid", cl.WorkHourController.DeleteWorkHourByUuid)

	//medicine with medical staff role
	medicine := echo.Group("api/v1/medicine")
	medicine.Use(middleware.JWTWithConfig(cl.JWTMiddleware), MedicalStaffValidation())
	medicine.POST("/", cl.MedicineController.CreateNewMedicine)
	medicine.PUT("/:uuid", cl.MedicineController.UpdateMedicineById)
	medicine.DELETE("/:uuid", cl.MedicineController.DeleteMedicineByUuid)

	//Medicine with doctor,medical staff and nurse role
	medicineWithAllRole := echo.Group("api/v1/medicine")
	medicineWithAllRole.Use(middleware.JWTWithConfig(cl.JWTMiddleware), AllRole())
	medicineWithAllRole.GET("/:uuid", cl.MedicineController.FindMedicineByUuid)
	medicineWithAllRole.GET("/queryName", cl.MedicineController.FindMedicineByNameQuery)
	medicineWithAllRole.GET("/", cl.MedicineController.GetMedicines)
	medicineWithAllRole.GET("/pagination", cl.MedicineController.GetMedicinePagination)

	//recipe with medical staff role
	recipe := echo.Group("api/v1/recipe")
	recipe.Use(middleware.JWTWithConfig(cl.JWTMiddleware), DoctorValidation())
	recipe.POST("/", cl.RecipeController.CreateNewRecipe)
	recipe.PUT("/:uuid", cl.RecipeController.UpdateRecipeById)
	recipe.DELETE("/:uuid", cl.RecipeController.DeleteRecipeByUuid)

	//recipe with doctor,medical staff and nurse role
	recipeWithAllRole := echo.Group("api/v1/recipe")
	recipeWithAllRole.Use(middleware.JWTWithConfig(cl.JWTMiddleware), AllRole())
	recipeWithAllRole.GET("/:uuid", cl.RecipeController.FindRecipeByUuid)
	recipeWithAllRole.GET("/:name", cl.RecipeController.FindRecipeByUuid)
	recipeWithAllRole.GET("/", cl.RecipeController.GetRecipe)

	//medical Record with doctor role
	medicalRecord := echo.Group("api/v1/medicalRecord")
	medicalRecord.Use(middleware.JWTWithConfig(cl.JWTMiddleware), AllRole())
	medicalRecord.POST("/", cl.MedicalRecordController.CreateMedicalRecord)
	medicalRecord.GET("/queue", cl.MedicalRecordController.GetMedicalRecordsQueue)
	medicalRecord.GET("/medicalStaff/queue", cl.MedicalRecordController.GetMedicalRecordsByMedicalStaff)
	medicalRecord.PUT("/:uuid", cl.MedicalRecordController.UpdateMedicalRecordById)
	medicalRecord.DELETE("/:uuid", cl.MedicalRecordController.DeleteMedicalRecordByUuid)
	medicalRecord.GET("/:uuid", cl.MedicalRecordController.FindMedicalRecordByUuid)
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
