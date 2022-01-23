package medicalRecordController

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalRecordEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalRecordController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalRecordController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type MedicalRecordsController struct {
	medicalRecordsService medicalRecordEntity.Service
	jwtAuth               *auth.ConfigJWT
}

func NewMedicalRecordsController(mrs medicalRecordEntity.Service, jwtAuth *auth.ConfigJWT) *MedicalRecordsController {
	return &MedicalRecordsController{
		medicalRecordsService: mrs,
		jwtAuth:               jwtAuth,
	}
}

func (ctrl *MedicalRecordsController) CreateMedicalRecord(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.MedicalRecordRegistration)

	patientID := c.QueryParam("patientID")
	userID := c.QueryParam("userID")
	medicalStaff := auth.GetDoctor(c)
	medicalStaffID := medicalStaff.Uuid

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

	res, err := ctrl.medicalRecordsService.CreateMedicalRecord(ctx, req.ToDomain(), userID, medicalStaffID, patientID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created Medical Record data",
			response.FromDomain(res)))
}

func (ctrl *MedicalRecordsController) FindMedicalRecordByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	medicalRecord, err := ctrl.medicalRecordsService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Record Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Medical Record By id",
			response.FromDomain(&medicalRecord)))
}

func (ctrl *MedicalRecordsController) GetMedicalRecordsQueue(c echo.Context) error {
	user := auth.GetUser(c)
	userID := user.Uuid

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	data, offset, limit, totalData, err := ctrl.medicalRecordsService.GetMedicalRecordsQueue(c.Request().Context(), userID, page)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Record Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.MedicalRecords{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(*data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("Successfully Get all Medical Record by nik But Medical Record  Data Doesn't Exist",
				data))
	}

	return helpers.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *MedicalRecordsController) UpdateMedicalRecordById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.MedicalRecordRegistration)

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
	patientID := c.QueryParam("patientID")
	userID := c.QueryParam("userID")
	medicalStaff := auth.GetDoctor(c)
	medicalStaffID := medicalStaff.Uuid

	res, err := ctrl.medicalRecordsService.UpdateById(ctx, req.ToDomain(), userID, medicalStaffID, patientID, uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an Medical Record",
			response.FromDomain(res)))
}

func (ctrl *MedicalRecordsController) DeleteMedicalRecordByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.medicalRecordsService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Record doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.medicalRecordsService.DeleteMedicalRecord(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Medical Record",
			nil))
}
