package recipe

import (
	"errors"

	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kdelwat/recipaliser/db"
	"github.com/kdelwat/recipaliser/ingredient"
)

type Recipe struct {
	gorm.Model
	Name        string
	Ingredients []ingredient.Ingredient `gorm:"many2many:recipe_ingredients;"`
}

func New(name string) (Recipe, error) {
	var newRecipe Recipe
	var existingRecipe Recipe

	err := db.Db.First(&existingRecipe, "name = ?", name).Error

	if err != gorm.ErrRecordNotFound {
		return newRecipe, errors.New("A recipe with the same name already exists")
	}

	err = db.Db.Create(&Recipe{Name: name, Ingredients: []ingredient.Ingredient{}}).Error

	if err != nil {
		return newRecipe, errors.New(fmt.Sprintf("Could not create new recipe in database: %v", err))
	}

	return newRecipe, nil
}
