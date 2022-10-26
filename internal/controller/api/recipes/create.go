package recipes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indeedhat/recipe-api/internal/repo"
)

type createRecipeRequest struct {
	Slug             string `binding:"required" json:"Slug"`
	Title            string `binding:"required" json:"Title"`
	Description      string `binding:"required" json:"Description"`
	ShortDescription string `binding:"required" json:"ShortDescription"`
	Ingredients      string `binding:"required" json:"Ingredients"`
	Steps            string `binding:"required" json:"Steps"`
}

// Create adds a new recipe to the database
func (c RecipeController) Create(ctx *gin.Context) {
	var (
		input       createRecipeRequest
		ingredients repo.Ingredients
		steps       repo.RecipeSteps
	)

	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Bad Input"))
		return
	}

	if err := ingredients.Scan(input.Ingredients); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Bad Ingredients"))
		return
	}

	if err := steps.Scan(input.Ingredients); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Bad Steps"))
		return
	}

	if existing, _ := c.recipeRepo.FindBySlug(ctx, input.Slug); existing != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Slug in use"))
		return
	}

	recipe := repo.Recipe{
		Slug:             input.Slug,
		Title:            input.Title,
		Description:      input.Description,
		ShortDescription: input.ShortDescription,
		Ingredients:      ingredients,
		Steps:            steps,
	}

	if err := c.recipeRepo.Create(ctx, &recipe); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Create failed"))
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}
