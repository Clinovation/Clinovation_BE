package medicalStaffController

import (
	"fmt"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/medicalStaffEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalStaffController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/medicalStaffController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"time"
)

type MedicalStaffController struct {
	medicalStaffService medicalStaffEntity.Service
	jwtAuth             *auth.ConfigJWT
}

func NewMedicalStaffController(mss medicalStaffEntity.Service, jwtAuth *auth.ConfigJWT) *MedicalStaffController {
	return &MedicalStaffController{
		medicalStaffService: mss,
		jwtAuth:             jwtAuth,
	}
}

func (ctrl *MedicalStaffController) Registration(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.MedicalStaffRegistration)

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

	res, err := ctrl.medicalStaffService.Register(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created medical staff account",
			response.FromDomain(res)))
}

func (ctrl *MedicalStaffController) LoginMedicalStaff(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.MedicalStaffLogin)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("The Data You Entered is Wrong",
				err, helpers.EmptyObj{}))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	token, err := ctrl.medicalStaffService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Staff Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	expire := time.Now().Add(1 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:    "is-login",
		Value:   token,
		Expires: expire,
	}
	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("successful to login",
			res))
}

func (ctrl *MedicalStaffController) FindMedicalStaffByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	medicalStaff, err := ctrl.medicalStaffService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Staff Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Medical Staff By id",
			response.FromDomain(&medicalStaff)))
}

func (ctrl *MedicalStaffController) FindMedicalStaffByNameQuery(c echo.Context) error {
	name := c.QueryParam("name")

	medicalStaff, err := ctrl.medicalStaffService.FindByName(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Staff Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Medical Staff  By Name",
			response.FromDomainArray(medicalStaff)))
}

func (ctrl *MedicalStaffController) FindMedicalStaffByNikQuery(c echo.Context) error {
	nik := c.QueryParam("nik")

	medicalStaff, err := ctrl.medicalStaffService.FindByNik(c.Request().Context(), nik)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Staff Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Medical Staff By Nik",
			response.FromDomainArray(medicalStaff)))
}

func (ctrl *MedicalStaffController) GetMedicalStaff(c echo.Context) error {
	medicalStaff, err := ctrl.medicalStaffService.GetMedicalStaff(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Staff Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Medical Staff",
			response.FromDomainArray(*medicalStaff)))
}

func (ctrl *MedicalStaffController) UpdateMedicalStaffById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.MedicalStaffRegistration)

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

	medicalStaff := auth.GetMedicalStaff(c)
	medicalStaffId := medicalStaff.Uuid

	res, err := ctrl.medicalStaffService.UpdateById(ctx, req.ToDomain(), medicalStaffId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an account",
			response.FromDomain(res)))
}

func (ctrl *MedicalStaffController) UploadAvatar(c echo.Context) error {
	ctx := c.Request().Context()
	//file := new(request.UserUploadAvatar)
	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	medicalStaff := auth.GetMedicalStaff(c)
	medicalStaffId := medicalStaff.Uuid

	path := fmt.Sprintf("images/avatar/%v-%s", medicalStaffId, file.Filename)

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

	res, err := ctrl.medicalStaffService.UploadAvatar(ctx, medicalStaffId, path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully upload avatar",
			response.FromDomain(res)))
}

func (ctrl *MedicalStaffController) DeleteMedicalStaffByUuid(c echo.Context) error {
	medicalStaff := auth.GetMedicalStaff(c)
	medicalStaffId := medicalStaff.Uuid

	_, errGet := ctrl.medicalStaffService.FindByUuid(c.Request().Context(), medicalStaffId)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medical Staff doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.medicalStaffService.DeleteMedicalStaff(c.Request().Context(), medicalStaffId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Medical Staff",
			nil))
}
