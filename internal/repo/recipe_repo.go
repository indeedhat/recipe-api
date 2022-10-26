package repo

import (
	"context"

	"gorm.io/gorm"
)

type RecipeRepo interface {
	FindById(ctx context.Context, id any) (*Recipe, error)
	FindBySlug(ctx context.Context, slug string) (*Recipe, error)
	Create(ctx context.Context, recipe *Recipe) error
	Update(ctx context.Context, recipe *Recipe) error
}

type RecipeSqlRepo struct {
	baseRepo
}

// Create adds a new recipe to the database
func (r RecipeSqlRepo) Create(ctx context.Context, recipe *Recipe) error {
	return r.db(ctx).Create(recipe).Error
}

// FindById finds a recipe by its unique id
func (r RecipeSqlRepo) FindById(ctx context.Context, id any) (*Recipe, error) {
	var recipe Recipe

	if err := r.db(ctx).First(&recipe, id).Error; err != nil {
		return nil, err
	}

	return &recipe, nil
}

// FindBySlug finds a recipe by its unique slug
func (r RecipeSqlRepo) FindBySlug(ctx context.Context, slug string) (*Recipe, error) {
	var recipe Recipe

	if err := r.db(ctx).First(&recipe, "slug = ?", slug).Error; err != nil {
		return nil, err
	}

	return &recipe, nil
}

// Update updates a recipe in the database
func (r RecipeSqlRepo) Update(ctx context.Context, recipe *Recipe) error {
	return r.db(ctx).Save(recipe).Error
}

// NewRecipeSqlRepo creates a new instance of the RecipeSqlRepo struct
func NewRecipeSqlRepo(fallbackDB *gorm.DB) RecipeRepo {
	return RecipeSqlRepo{
		baseRepo{fallbackDB: fallbackDB},
	}
}
