package workDayController

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/workDayEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/workDayController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/workDayController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type WorkDayController struct {
	workDaysService workDayEntity.Service
	jwtAuth         *auth.ConfigJWT
}

func NewWorkDayController(wds workDayEntity.Service, jwtAuth *auth.ConfigJWT) *WorkDayController {
	return &WorkDayController{
		workDaysService: wds,
		jwtAuth:         jwtAuth,
	}
}

func (ctrl *WorkDayController) CreateNewWorkDay(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.WorkDayRegistration)

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

	res, err := ctrl.workDaysService.CreateWorkDay(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Created Work Day",
			response.FromDomain(res)))
}

func (ctrl *WorkDayController) FindWorkDayByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	workDay, err := ctrl.workDaysService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("work Day Doesn't Exist By id",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get work Day By id",
			response.FromDomain(&workDay)))
}

func (ctrl *WorkDayController) FindWorkDayByDay(c echo.Context) error {
	day := c.QueryParam("day")

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	data, offset, limit, totalData, err := ctrl.workDaysService.FindByDay(c.Request().Context(), day, page)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Work Day Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.WorkDays{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("Successfully Get all Work Day by day But Work Day Data Doesn't Exist",
				data))
	}

	return helpers.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *WorkDayController) GetWorkDaysPagination(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.workDaysService.GetWorkDaysPagination(c.Request().Context(), page)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Work Day Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	res := []response.WorkDays{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(*data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("Successfully Get all Work Day But Work Day Data Doesn't Exist",
				data))
	}

	return helpers.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *WorkDayController) UpdateWorkDayById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.WorkDayRegistration)

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

	res, err := ctrl.workDaysService.UpdateById(ctx, req.ToDomain(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an account",
			response.FromDomain(res)))
}

func (ctrl *WorkDayController) DeleteWorkDayByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.workDaysService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Work Day doesn't exist by delete",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.workDaysService.DeleteWorkDay(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Work Day",
			nil))
}

func (ctrl *WorkDayController) GetWorkDays(c echo.Context) error {
	recipe, err := ctrl.workDaysService.GetWorkDays(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Work Days Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Work Days",
			response.FromDomainArray(*recipe)))
}
