package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/indeedhat/juniper"
	"github.com/indeedhat/recipe-api/internal/controller/api/recipes"
	"github.com/indeedhat/recipe-api/internal/repo"
)

func main() {
	db, err := repo.Connect()
	if err != nil {
		log.Fatalf("db error: %s", err)
	}

	if err := repo.Migrate(db); err != nil {
		log.Fatalf("db migration error: %s", err)
	}

	recipeRepo := repo.NewRecipeSqlRepo(db)

	router := gin.Default()
	router.Use(juniper.DBTransactionMiddleware(db))

	_ = recipes.New(router, recipeRepo)

	router.Run()
}
