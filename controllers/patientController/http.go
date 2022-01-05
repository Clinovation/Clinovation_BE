package patientController

import (
	"fmt"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/patientEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/patientController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/patientController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

type PatientsController struct {
	patientsService patientEntity.Service
	jwtAuth         *auth.ConfigJWT
}

func NewPatientsController(ps patientEntity.Service, jwtAuth *auth.ConfigJWT) *PatientsController {
	return &PatientsController{
		patientsService: ps,
		jwtAuth:         jwtAuth,
	}
}

func (ctrl *PatientsController) Registration(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.PatientRegistration)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("The Data You Entered is Wrong",
				err, helpers.EmptyObj{}))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An Error Occurred While Validating The Request Data",
				err, helpers.EmptyObj{}))
	}

	medicalStaff := auth.GetMedicalStaff(c)
	medicalStaffId := medicalStaff.Uuid

	res, err := ctrl.patientsService.Register(ctx, req.ToDomain(), medicalStaffId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created patient data",
			response.FromDomain(res)))
}

func (ctrl *PatientsController) FindPatientByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	patient, err := ctrl.patientsService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Patient Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Patient By id",
			response.FromDomain(&patient)))
}

func (ctrl *PatientsController) FindPatientByNameQuery(c echo.Context) error {
	name := c.QueryParam("name")

	patients, err := ctrl.patientsService.FindByName(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Patient Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Patient By Name",
			response.FromDomainArray(patients)))
}

func (ctrl *PatientsController) FindPatientByNikQuery(c echo.Context) error {
	nik := c.QueryParam("nik")

	patients, err := ctrl.patientsService.FindByNik(c.Request().Context(), nik)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Patient Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Patient By Nik",
			response.FromDomainArray(patients)))
}

func (ctrl *PatientsController) GetPatients(c echo.Context) error {
	patients, err := ctrl.patientsService.GetPatients(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Patient Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Patient",
			response.FromDomainArray(*patients)))
}

func (ctrl *PatientsController) UpdatePatientById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.PatientRegistration)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while input the data",
				err, helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	uuid := c.Param("uuid")

	res, err := ctrl.patientsService.UpdateById(ctx, req.ToDomain(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an account",
			response.FromDomain(res)))
}

func (ctrl *PatientsController) UploadAvatar(c echo.Context) error {
	ctx := c.Request().Context()
	//file := new(request.UserUploadAvatar)
	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	uuid := c.Param("uuid")

	path := fmt.Sprintf("images/avatar/%v-%s", uuid, file.Filename)

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}
	defer src.Close()

	destination, err := os.Create(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}
	defer destination.Close()

	if _, err = io.Copy(destination, src); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.patientsService.UploadAvatar(ctx, uuid, path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully upload avatar",
			response.FromDomain(res)))
}

func (ctrl *PatientsController) DeletePatientByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.patientsService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Patient doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.patientsService.DeletePatient(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Patient",
			nil))
}
