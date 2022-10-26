package repo

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/indeedhat/juniper"
)

// Recipe is the database model for recipes
type Recipe struct {
	juniper.Model

	Slug             string `gorm:"uniqueIndex"`
	Title            string
	Description      string
	ShortDescription string
	Image            string

	Ingredients Ingredients
	Steps       RecipeSteps
}

type Ingredients map[string]Ingredient

// Scan implements sql.Scanner
func (i *Ingredients) Scan(src any) error {
	switch value := src.(type) {
	case []byte:
		return json.Unmarshal(value, i)
	case string:
		return json.Unmarshal([]byte(value), i)
	default:
		return errors.New("invalid type")
	}
}

// Value implements driver.Valuer
func (i *Ingredients) Value() (driver.Value, error) {
	val, err := json.Marshal(i)
	return string(val), err

}

var _ sql.Scanner = (*Ingredients)(nil)
var _ driver.Valuer = (*Ingredients)(nil)

type Ingredient struct {
	Amount string
	Text   string
}

type RecipeSteps []RecipeStep

// Scan implements sql.Scanner
func (r *RecipeSteps) Scan(src any) error {
	switch value := src.(type) {
	case []byte:
		return json.Unmarshal(value, r)
	case string:
		return json.Unmarshal([]byte(value), r)
	default:
		return errors.New("invalid type")
	}
}

// Value implements driver.Valuer
func (r *RecipeSteps) Value() (driver.Value, error) {
	val, err := json.Marshal(r)
	return string(val), err

}

var _ sql.Scanner = (*RecipeSteps)(nil)
var _ driver.Valuer = (*RecipeSteps)(nil)

type RecipeStep struct {
	Ingredients []string
	Time        time.Duration
	Text        string
}
