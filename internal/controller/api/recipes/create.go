package recipes

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/indeedhat/recipe-api/internal/repo"
)

type createRecipeRequest struct {
	Slug        string           `binding:"required" json:"Slug"`
	Title       string           `binding:"required" json:"Title"`
	Description string           `binding:"required" json:"Description"`
	CookTime    time.Duration    `binding:"required" json:"CookTime"`
	PrepTime    time.Duration    `binding:"required" json:"PrepTime"`
	Ingredients repo.Ingredients `binding:"required" json:"Ingredients"`
	Steps       repo.RecipeSteps `binding:"required" json:"Steps"`
}

// Create adds a new recipe to the database
func (c RecipeController) Create(ctx *gin.Context) {
	var input createRecipeRequest

	if err := ctx.BindJSON(&input); err != nil {
		if val, ok := err.(validator.ValidationErrors); ok {
			for _, e := range val {
				log.Printf("%s: %s", e.Field(), e.Error())
			}
		} else {
			log.Printf("not a validation error: %s", err)
		}
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Bad Input"))
		return
	}

	if existing, _ := c.recipeRepo.FindBySlug(ctx, input.Slug); existing != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Slug in use"))
		return
	}

	// TODO: image

	recipe := repo.Recipe{
		Slug:        input.Slug,
		Title:       input.Title,
		Description: input.Description,
		Ingredients: input.Ingredients,
		Steps:       input.Steps,
		CookTime:    input.CookTime,
		PrepTime:    input.PrepTime,
	}

	spew.Dump(recipe)

	if err := c.recipeRepo.Create(ctx, &recipe); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Create failed"))
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}
