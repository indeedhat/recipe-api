package recipes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/indeedhat/recipe-api/internal/repo"
)

type updateRecipeRequest struct {
	Title       string           `binding:"required" json:"Title"`
	Description string           `binding:"required" json:"Description"`
	CookTime    time.Duration    `binding:"required" json:"CookTime"`
	PrepTime    time.Duration    `binding:"required" json:"PrepTime"`
	Ingredients repo.Ingredients `binding:"required" json:"Ingredients"`
	Steps       repo.RecipeSteps `binding:"required" json:"Steps"`
}

// Update updates an existing recipe in the database
func (c RecipeController) Update(ctx *gin.Context) {
	var (
		input updateRecipeRequest
		slug  = ctx.Param("slug")
	)

	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Bad Input"))
		return
	}

	existing, _ := c.recipeRepo.FindBySlug(ctx, slug)
	if existing != nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("Recipe not found"))
		return
	}

	existing.Title = input.Title
	existing.Description = input.Description
	existing.CookTime = input.CookTime
	existing.PrepTime = input.PrepTime
	existing.Ingredients = input.Ingredients
	existing.Steps = input.Steps

	if err := c.recipeRepo.Create(ctx, existing); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Create failed"))
		return
	}

	ctx.JSON(http.StatusOK, existing)
}
