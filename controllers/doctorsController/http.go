package doctorsController

import (
	"fmt"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"time"
)

type DoctorController struct {
	doctorsService doctorsEntity.Service
	jwtAuth        *auth.ConfigJWT
}

func NewDoctorController(de doctorsEntity.Service, jwtAuth *auth.ConfigJWT) *DoctorController {
	return &DoctorController{
		doctorsService: de,
		jwtAuth:        jwtAuth,
	}
}

func (ctrl *DoctorController) Registration(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.DoctorRegistration)

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

	res, err := ctrl.doctorsService.Register(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created doctor account",
			response.FromDomain(res)))
}

func (ctrl *DoctorController) LoginDoctor(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.DoctorLogin)

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

	token, err := ctrl.doctorsService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Doesn't Exist",
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

func (ctrl *DoctorController) FindDoctorByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	doctor, err := ctrl.doctorsService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get doctor By id",
			response.FromDomain(&doctor)))
}

func (ctrl *DoctorController) GetDoctors(c echo.Context) error {
	doctor, err := ctrl.doctorsService.GetDoctors(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all doctors",
			response.FromDomainArray(*doctor)))
}

func (ctrl *DoctorController) UpdateDoctorById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.DoctorRegistration)

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

	doctor := auth.GetDoctor(c)
	doctorId := doctor.Uuid

	res, err := ctrl.doctorsService.UpdateById(ctx, req.ToDomain(), doctorId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an account",
			response.FromDomain(res)))
}

func (ctrl *DoctorController) UploadAvatar(c echo.Context) error {
	ctx := c.Request().Context()
	//file := new(request.UserUploadAvatar)
	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Upload Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	doctor := auth.GetDoctor(c)
	doctorId := doctor.Uuid

	path := fmt.Sprintf("images/avatar/%v-%s", doctorId, file.Filename)

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

	res, err := ctrl.doctorsService.UploadAvatar(ctx, doctorId, path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully upload avatar",
			response.FromDomain(res)))
}

func (ctrl *DoctorController) DeleteDoctorByUuid(c echo.Context) error {
	doctor := auth.GetDoctor(c)
	doctorId := doctor.Uuid

	_, errGet := ctrl.doctorsService.FindByUuid(c.Request().Context(), doctorId)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.doctorsService.DeleteDoctor(c.Request().Context(), doctorId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Doctor",
			nil))
}
