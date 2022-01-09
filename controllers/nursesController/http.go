package nursesController

import (
	"fmt"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/nursesEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/nursesController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/nursesController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"time"
)

type NurseController struct {
	nursesService nursesEntity.Service
	jwtAuth       *auth.ConfigJWT
}

func NewNursesController(ns nursesEntity.Service, jwtAuth *auth.ConfigJWT) *NurseController {
	return &NurseController{
		nursesService: ns,
		jwtAuth:       jwtAuth,
	}
}

func (ctrl *NurseController) Registration(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.NurseRegistration)

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

	res, err := ctrl.nursesService.Register(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created nurse account",
			response.FromDomain(res)))
}

func (ctrl *NurseController) LoginNurse(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.NurseLogin)

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

	token, err := ctrl.nursesService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Doesn't Exist",
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

func (ctrl *NurseController) FindNurseByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	nurse, err := ctrl.nursesService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Nurse By id",
			response.FromDomain(&nurse)))
}

func (ctrl *NurseController) FindNurseByNameQuery(c echo.Context) error {
	name := c.QueryParam("name")

	nurse, err := ctrl.nursesService.FindByName(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Nurse By Name",
			response.FromDomainArray(nurse)))
}

func (ctrl *NurseController) FindNurseByNikQuery(c echo.Context) error {
	nik := c.QueryParam("nik")

	nurse, err := ctrl.nursesService.FindByNik(c.Request().Context(), nik)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Nurse By Nik",
			response.FromDomainArray(nurse)))
}

func (ctrl *NurseController) GetNurses(c echo.Context) error {
	nurse, err := ctrl.nursesService.GetNurses(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Nurse",
			response.FromDomainArray(*nurse)))
}

func (ctrl *NurseController) UpdateNurseById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.NurseRegistration)

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

	nurse := auth.GetNurse(c)
	nurseId := nurse.Uuid

	res, err := ctrl.nursesService.UpdateById(ctx, req.ToDomain(), nurseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an account",
			response.FromDomain(res)))
}

func (ctrl *NurseController) UploadAvatar(c echo.Context) error {
	ctx := c.Request().Context()
	//file := new(request.UserUploadAvatar)
	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	nurse := auth.GetNurse(c)
	nurseId := nurse.Uuid

	path := fmt.Sprintf("images/avatar/%v-%s", nurseId, file.Filename)

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

	res, err := ctrl.nursesService.UploadAvatar(ctx, nurseId, path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully upload avatar",
			response.FromDomain(res)))
}

func (ctrl *NurseController) DeleteNurseByUuid(c echo.Context) error {
	nurse := auth.GetNurse(c)
	nurseId := nurse.Uuid

	_, errGet := ctrl.nursesService.FindByUuid(c.Request().Context(), nurseId)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.nursesService.DeleteNurse(c.Request().Context(), nurseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Nurse",
			nil))
}