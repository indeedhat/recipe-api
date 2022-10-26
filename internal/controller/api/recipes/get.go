package recipes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get finds a recipe by its slug
func (c RecipeController) Get(ctx *gin.Context) {
	slug := ctx.Param("slug")

	existing, _ := c.recipeRepo.FindBySlug(ctx, slug)
	if existing != nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("Recipe not found"))
		return
	}

	ctx.JSON(http.StatusOK, existing)
}
