package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/indeedhat/recipe-api/internal/repo"
)

type RecipeController struct {
	recipeRepo repo.RecipeRepo
}

// New sets up and returns a new RecipeController
func New(router *gin.Engine, recipeRepo repo.RecipeRepo) RecipeController {
	controller := RecipeController{recipeRepo: recipeRepo}

	controller.route(router)

	return controller
}

func (c RecipeController) route(router *gin.Engine) {
	group := router.Group("/api/recipe")
	{
		group.GET(":slug", c.Get)
		group.POST("", c.Create)
		group.POST(":slug", c.Update)
	}
}
