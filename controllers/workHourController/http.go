package workHourController

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/workHourEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/workHourController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/workHourController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type WorkHourController struct {
	workHoursService workHourEntity.Service
	jwtAuth          *auth.ConfigJWT
}

func NewWorkHourController(wds workHourEntity.Service, jwtAuth *auth.ConfigJWT) *WorkHourController {
	return &WorkHourController{
		workHoursService: wds,
		jwtAuth:          jwtAuth,
	}
}

func (ctrl *WorkHourController) CreateNewWorkHour(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.WorkHourRegistration)

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

	res, err := ctrl.workHoursService.CreateWorkHour(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Created Work Hour",
			response.FromDomain(res)))
}

func (ctrl *WorkHourController) FindWorkHourByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	workHour, err := ctrl.workHoursService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("work Hour Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get work Hour By id",
			response.FromDomain(&workHour)))
}

func (ctrl *WorkHourController) FindWorkHourByHour(c echo.Context) error {
	hour := c.Param("hour")

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	data, offset, limit, totalData, err := ctrl.workHoursService.FindByHour(c.Request().Context(), hour, page)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Work Hour Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.WorkHours{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("Successfully Get all Work Hour by nik But Work Hour Data Doesn't Exist",
				data))
	}

	return helpers.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *WorkHourController) GetWorkHours(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	data, offset, limit, totalData, err := ctrl.workHoursService.GetWorkHours(c.Request().Context(), page)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Work Hour Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.WorkHours{}
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

func (ctrl *WorkHourController) UpdateWorkHourById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.WorkHourRegistration)

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

	res, err := ctrl.workHoursService.UpdateById(ctx, req.ToDomain(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an account",
			response.FromDomain(res)))
}

func (ctrl *WorkHourController) DeleteWorkHourByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.workHoursService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Work Hour doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.workHoursService.DeleteWorkHour(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Work Hour",
			nil))
}
