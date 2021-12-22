package doctorsController

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"fmt"
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

func NewDoctorController(de doctorsEntity.Service, jwtauth *auth.ConfigJWT) *DoctorController {
	return &DoctorController{
		doctorsService: de,
		jwtAuth:        jwtauth,
	}
}

func (ctrl *DoctorController) Registration(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.DoctorRegistration)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.doctorsService.Register(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created an account",
			response.FromDomain(res)))
}

func (ctrl *DoctorController) LoginDoctor(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.DoctorLogin)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	token, err := ctrl.doctorsService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	expire := time.Now().Add(5 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:    "is-login",
		Value:   token,
		Expires: expire,
	}
	c.SetCookie(&cookie)

	//return helpers.BuildSuccessResponseContext(c, response)
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("successful to login",
			response))
}

func (ctrl *DoctorController) LogoutDoctor(c echo.Context) error {
	err := ctrl.doctorsService.Logout(c)
	if err != nil {
		return c.JSON(http.StatusRequestTimeout,
			helpers.BuildSuccessResponse("you are not logged in",
				nil))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("you have successfully logged out",
			nil))
}

func (ctrl *DoctorController) FindDoctorByUuid(c echo.Context) error {
	//checking cookie is the doctor was login or not
	_, errC := c.Cookie("is-login")
	if errC != nil {
		return c.JSON(http.StatusForbidden,
			helpers.BuildErrorResponse("you are not logged in",
				errC, helpers.EmptyObj{}))
	}

	uuid := c.Param("uuid")

	doctor, err := ctrl.doctorsService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("doctor doesn't exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get doctor By id",
			response.FromDomain(&doctor)))
}

func (ctrl *DoctorController) UpdateDoctorById(c echo.Context) error {
	_, errC := c.Cookie("is-login")
	if errC != nil {
		return c.JSON(http.StatusForbidden,
			helpers.BuildErrorResponse("you are not logged in",
				errC, helpers.EmptyObj{}))
	}

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

	doctor := auth.GetUser(c)
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
	_, errC := c.Cookie("is-login")
	if errC != nil {
		return c.JSON(http.StatusForbidden,
			helpers.BuildErrorResponse("you are not logged in",
				errC, helpers.EmptyObj{}))
	}
	//var err error
	ctx := c.Request().Context()
	//file := new(request.UserUploadAvatar)
	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	doctor := auth.GetUser(c)
	doctorId := doctor.Uuid

	path := fmt.Sprintf("images/avatar/%v-%s", doctorId, file.Filename)

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}
	defer src.Close()

	destination, err := os.Create(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}
	defer destination.Close()

	if _, err = io.Copy(destination, src); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
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
	//checking login doctor
	_, errC := c.Cookie("is-login")
	if errC != nil {
		return c.JSON(http.StatusForbidden,
			helpers.BuildErrorResponse("you are not logged in",
				errC, helpers.EmptyObj{}))
	}

	doctor := auth.GetUser(c)
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
