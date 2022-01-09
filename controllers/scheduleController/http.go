package scheduleController

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/scheduleEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/scheduleController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/scheduleController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SchedulesController struct {
	schedulesService scheduleEntity.Service
	jwtAuth          *auth.ConfigJWT
}

func NewSchedulesController(ps scheduleEntity.Service, jwtAuth *auth.ConfigJWT) *SchedulesController {
	return &SchedulesController{
		schedulesService: ps,
		jwtAuth:          jwtAuth,
	}
}

func (ctrl *SchedulesController) CreateSchedule(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.ScheduleRegistration)

	uuid := c.Param("uuid")

	workDayID := c.QueryParam("workDayId")
	workHourID := c.QueryParam("workHourId")

	res, err := ctrl.schedulesService.CreateSchedule(ctx, req.ToDomain(), uuid, workDayID, workHourID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created Schedule data",
			response.FromDomain(res)))
}

func (ctrl *SchedulesController) FindScheduleByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	schedule, err := ctrl.schedulesService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Schedule Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Schedule By id",
			response.FromDomain(&schedule)))
}

func (ctrl *SchedulesController) GetDoctorSchedules(c echo.Context) error {
	schedules, err := ctrl.schedulesService.GetDoctorSchedules(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Schedule Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Doctor Schedule",
			response.FromDomainArray(*schedules)))
}

func (ctrl *SchedulesController) GetNurseSchedules(c echo.Context) error {
	schedules, err := ctrl.schedulesService.GetNurseSchedules(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Schedules Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Nurse Schedules",
			response.FromDomainArray(*schedules)))
}

func (ctrl *SchedulesController) UpdateScheduleById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.ScheduleRegistration)

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
	user := auth.GetUser(c)
	userID := user.Uuid

	workDayID := c.QueryParam("workDayId")
	workHourID := c.QueryParam("workHourId")

	res, err := ctrl.schedulesService.UpdateById(ctx, req.ToDomain(), userID, workDayID, workHourID, uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an schedule",
			response.FromDomain(res)))
}

func (ctrl *SchedulesController) DeleteScheduleByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.schedulesService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Schedule doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.schedulesService.DeleteSchedule(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Schedule",
			nil))
}

func (ctrl *SchedulesController) GetDoctorSchedulesByDay(c echo.Context) error {
	day := c.QueryParam("day")

	schedules, err := ctrl.schedulesService.GetDoctorSchedulesByDay(c.Request().Context(), day)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Schedules Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Doctor Schedules By Day",
			response.FromDomainArray(*schedules)))
}

func (ctrl *SchedulesController) GetDoctorSchedulesByHour(c echo.Context) error {
	hour := c.QueryParam("hour")

	schedules, err := ctrl.schedulesService.GetDoctorSchedulesByDay(c.Request().Context(), hour)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Schedules Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Doctor Schedules By Hour",
			response.FromDomainArray(*schedules)))
}

func (ctrl *SchedulesController) GetNurseSchedulesByHour(c echo.Context) error {
	hour := c.QueryParam("hour")

	Schedules, err := ctrl.schedulesService.GetNurseSchedulesByHour(c.Request().Context(), hour)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Schedule Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Nurse Schedule By Hour",
			response.FromDomainArray(*Schedules)))
}

func (ctrl *SchedulesController) GetNurseSchedulesByDay(c echo.Context) error {
	day := c.QueryParam("day")

	Schedules, err := ctrl.schedulesService.GetNurseSchedulesByDay(c.Request().Context(), day)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Schedule Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Nurse Schedule By Day",
			response.FromDomainArray(*Schedules)))
}
