package medicineController

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/medicineEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/medicineController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/medicineController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MedicineController struct {
	medicineService medicineEntity.Service
	jwtAuth         *auth.ConfigJWT
}

func NewMedicineController(ms medicineEntity.Service, jwtAuth *auth.ConfigJWT) *MedicineController {
	return &MedicineController{
		medicineService: ms,
		jwtAuth:         jwtAuth,
	}
}

func (ctrl *MedicineController) CreateNewMedicine(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.MedicineRegistration)

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

	res, err := ctrl.medicineService.CreateNewMedicine(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created medical staff account",
			response.FromDomain(res)))
}

func (ctrl *MedicineController) FindMedicineByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	medicine, err := ctrl.medicineService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medicine Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Medicine By id",
			response.FromDomain(&medicine)))
}

func (ctrl *MedicineController) FindMedicineByNameQuery(c echo.Context) error {
	name := c.QueryParam("name")

	medicine, err := ctrl.medicineService.FindByName(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medicine Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Medicine  By Name",
			response.FromDomainArray(medicine)))
}

func (ctrl *MedicineController) GetMedicine(c echo.Context) error {
	medicine, err := ctrl.medicineService.GetMedicines(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medicine Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Medicine",
			response.FromDomainArray(*medicine)))
}

func (ctrl *MedicineController) UpdateMedicineById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.MedicineRegistration)

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

	res, err := ctrl.medicineService.UpdateById(ctx, req.ToDomain(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an medicine",
			response.FromDomain(res)))
}

func (ctrl *MedicineController) DeleteMedicineByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.medicineService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Medicine doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.medicineService.DeleteMedicine(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Medicine",
			nil))
}
