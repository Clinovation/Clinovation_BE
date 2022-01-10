package recipeEntity

import (
	"context"
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"github.com/Clinovation/Clinovation_BE/businesses"
	"time"
)

type RecipeServices struct {
	recipesRepository Repository
	jwtAuth           *auth.ConfigJWT
	ContextTimeout    time.Duration
}

func NewRecipeServices(repoRecipe Repository, auth *auth.ConfigJWT, timeout time.Duration) Service {
	return &RecipeServices{
		recipesRepository: repoRecipe,
		jwtAuth:           auth,
		ContextTimeout:    timeout,
	}
}

func (qs *RecipeServices) CreateNewRecipe(ctx context.Context, recipeDomain *Domain) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	res, err := qs.recipesRepository.CreateNewRecipe(ctx, recipeDomain)
	if err != nil {
		return nil, businesses.ErrInternalServer
	}

	return res, nil
}

func (qs *RecipeServices) FindByUuid(ctx context.Context, uuid string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	result, err := qs.recipesRepository.GetByUuid(ctx, uuid)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (qs *RecipeServices) UpdateById(ctx context.Context, recipeDomain *Domain, id string) (*Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	result, err := qs.recipesRepository.UpdateRecipe(ctx, id, recipeDomain)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (qs *RecipeServices) DeleteRecipe(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	res, err := qs.recipesRepository.DeleteRecipeByUuid(ctx, id)
	if err != nil {
		return "", businesses.ErrNotFoundRecipe
	}
	return res, nil
}

func (qs *RecipeServices) GetRecipes(ctx context.Context) (*[]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, qs.ContextTimeout)
	defer cancel()

	res, err := qs.recipesRepository.GetRecipe(ctx)
	if err != nil {
		return &[]Domain{}, err
	}
	return res, nil
}

