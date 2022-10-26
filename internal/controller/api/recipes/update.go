package recipes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indeedhat/recipe-api/internal/repo"
)

type updateRecipeRequest struct {
	Title            string `binding:"required" json:"Title"`
	Description      string `binding:"required" json:"Description"`
	ShortDescription string `binding:"required" json:"ShortDescription"`
	Ingredients      string `binding:"required" json:"Ingredients"`
	Steps            string `binding:"required" json:"Steps"`
}

// Update updates an existing recipe in the database
func (c RecipeController) Update(ctx *gin.Context) {
	var (
		input       updateRecipeRequest
		ingredients repo.Ingredients
		steps       repo.RecipeSteps
		slug        = ctx.Param("slug")
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

	existing, _ := c.recipeRepo.FindBySlug(ctx, slug)
	if existing != nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("Recipe not found"))
		return
	}

	existing.Title = input.Title
	existing.Description = input.Description
	existing.ShortDescription = input.ShortDescription
	existing.Ingredients = ingredients
	existing.Steps = steps

	if err := c.recipeRepo.Create(ctx, existing); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Create failed"))
		return
	}

	ctx.JSON(http.StatusOK, existing)
}
