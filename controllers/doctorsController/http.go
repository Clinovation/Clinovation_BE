package doctorsController

import (
	"fmt"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/doctorsEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/doctorsController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strconv"
)

type DoctorController struct {
	doctorsService doctorsEntity.Service
	jwtAuth        *auth.ConfigJWT
}

func NewDoctorController(ds doctorsEntity.Service, jwtAuth *auth.ConfigJWT) *DoctorController {
	return &DoctorController{
		doctorsService: ds,
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

	//expire := time.Now().Add(1 * 24 * time.Hour)
	//cookie := http.Cookie{
	//	Name:    "is-login",
	//	Value:   token,
	//	Expires: expire,
	//}
	//c.SetCookie(&cookie)

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

func (ctrl *DoctorController) FindByJwt(c echo.Context) error {
	doctor := auth.GetDoctor(c)
	doctorId := doctor.Uuid

	res, err := ctrl.doctorsService.FindByUuid(c.Request().Context(), doctorId)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get doctor By id With Jwt",
			response.FromDomain(&res)))
}

func (ctrl *DoctorController) AcceptDoctor(c echo.Context) error {
	uuid := c.Param("uuid")

	doctor, err := ctrl.doctorsService.AcceptDoctor(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Accept doctor By id",
			response.FromDomain(doctor)))
}

func (ctrl *DoctorController) FindDoctorByNameQuery(c echo.Context) error {
	name := c.QueryParam("name")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	data, offset, limit, totalData, err := ctrl.doctorsService.FindByName(c.Request().Context(), name, page)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.Doctors{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)
	if len(data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("Successfully Get all Doctors by name But Doctor Data Doesn't Exist",
				data))
	}

	return helpers.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *DoctorController) FindDoctorByNikQuery(c echo.Context) error {
	nik := c.QueryParam("nik")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	data, offset, limit, totalData, err := ctrl.doctorsService.FindByNik(c.Request().Context(), nik, page)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Patient Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.Doctors{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("Successfully Get all Doctors by nik But Doctor  Data Doesn't Exist",
				data))
	}

	return helpers.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *DoctorController) GetWaitingList(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	data, offset, limit, totalData, err := ctrl.doctorsService.GetWaitingList(c.Request().Context(), page)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.Doctors{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(*data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("Successfully Get all Doctors Waiting List But Doctor Data Doesn't Exist",
				data))
	}

	return helpers.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *DoctorController) GetDoctors(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.doctorsService.GetDoctors(c.Request().Context(), page)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.Doctors{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(*data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("Successfully Get all Doctors But Doctor Data Doesn't Exist",
				data))
	}

	return helpers.NewSuccessResponse(c, http.StatusOK, res, resPage)
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

func (ctrl *DoctorController) ChangePassword(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.ChangePassword)

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

	res, err := ctrl.doctorsService.ChangePassword(ctx, req.ToDomainChange(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an account",
			response.FromDomain(res)))
}

func (ctrl *DoctorController) ForgetPassword(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.ForgetPassword)

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

	res, err := ctrl.doctorsService.ForgetPassword(ctx, req.ToDomainForget())
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

func (ctrl *DoctorController) DeleteDoctorByMedicalStaff(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.doctorsService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.doctorsService.DeleteDoctor(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Doctor",
			nil))
}
