package recipeController

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses/recipeEntity"
	"github.com/Clinovation/Clinovation_BE/controllers/recipeController/request"
	"github.com/Clinovation/Clinovation_BE/controllers/recipeController/response"
	"github.com/Clinovation/Clinovation_BE/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RecipeController struct {
	recipeService recipeEntity.Service
	jwtAuth       *auth.ConfigJWT
}

func NewRecipeController(ms recipeEntity.Service, jwtAuth *auth.ConfigJWT) *RecipeController {
	return &RecipeController{
		recipeService: ms,
		jwtAuth:       jwtAuth,
	}
}

func (ctrl *RecipeController) CreateNewRecipe(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.RecipeRegistration)

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

	res, err := ctrl.recipeService.CreateNewRecipe(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created Recipe",
			response.FromDomain(res)))
}

func (ctrl *RecipeController) FindRecipeByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	recipe, err := ctrl.recipeService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Recipe Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Recipe By id",
			response.FromDomain(&recipe)))
}

func (ctrl *RecipeController) GetRecipe(c echo.Context) error {
	recipe, err := ctrl.recipeService.GetRecipes(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Recipe Doesn't Exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get all Recipe",
			response.FromDomainArray(*recipe)))
}

func (ctrl *RecipeController) UpdateRecipeById(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.RecipeRegistration)

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

	res, err := ctrl.recipeService.UpdateById(ctx, req.ToDomain(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an Recipe",
			response.FromDomain(res)))
}

func (ctrl *RecipeController) DeleteRecipeByUuid(c echo.Context) error {
	uuid := c.Param("uuid")

	_, errGet := ctrl.recipeService.FindByUuid(c.Request().Context(), uuid)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Recipe doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.recipeService.DeleteRecipe(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a Recipe",
			nil))
}
