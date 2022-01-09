package queueController

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/queueEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/queueController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/queueController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type QueuesController struct {
	queuesService queueEntity.Service
	jwtAuth       *auth.ConfigJWT
}

func NewQueuesController(qs queueEntity.Service, jwtAuth *auth.ConfigJWT) *QueuesController {
	return &QueuesController{
		queuesService: qs,
		jwtAuth:       jwtAuth,
	}
}

func (ctrl *QueuesController) CreateQueue(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.QueueRegistration)

	uuid := c.Param("uuid")

	workDayID := c.QueryParam("workDayId")
	workHourID := c.QueryParam("workHourId")

	res, err := ctrl.queuesService.CreateQueue(ctx, req.ToDomain(), uuid, workDayID, workHourID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created Queue data",
			response.FromDomain(res)))
}

func (ctrl *QueuesController) FindQueueByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	queue, err := ctrl.queuesService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Queue Data Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Queue By id",
			response.FromDomain(&queue)))
}

func (ctrl *QueuesController) GetDoctorQueues(c echo.Context) error {
	queues, err := ctrl.queuesService.GetDoctorQueues(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Doctor Queue Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Doctor Queue",
			response.FromDomainArray(*queues)))
}

func (ctrl *QueuesController) GetNurseQueues(c echo.Context) error {
	queues, err := ctrl.queuesService.GetNurseQueues(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Nurse Queue Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Nurse Queue",
			response.FromDomainArray(*queues)))
}

func (ctrl *QueuesController) UpdateQueueById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.QueueRegistration)

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

	userID := c.QueryParam("userID")
	workDayID := c.QueryParam("workDayId")
	workHourID := c.QueryParam("workHourId")

	res, err := ctrl.queuesService.UpdateById(ctx, req.ToDomain(), userID, workDayID, workHourID, uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an Queue",
			response.FromDomain(res)))
}

func (ctrl *QueuesController) DeleteQueueByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.queuesService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Queue doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.queuesService.DeleteQueue(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Queue",
			nil))
}
