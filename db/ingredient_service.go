package db

import (
	"github.com/kdelwat/recipaliser"
)

var _ recipaliser.IngredientService = &IngredientService{}

type IngredientService struct {
	database *Database
}

func (is *IngredientService) Ingredient(id recipaliser.IngredientID) (*recipaliser.Ingredient, error) {
	return nil, nil
}

func (is *IngredientService) CreateIngredient(ingredient *recipaliser.Ingredient) error {
	var existingIngredient recipaliser.Ingredient

	if err := is.database.Collection("ingredients").Find("name = ?", ingredient.Name).One(&existingIngredient); err != nil {
		if err.Error() != "upper: no more rows in this result set" {
			return err
		}
	} else {
		return recipaliser.IngredientAlreadyExists
	}

	_, err := is.database.Collection("ingredients").Insert(*ingredient)

	return err
}

func (is *IngredientService) SearchIngredient(nameSubstring string) ([]*recipaliser.Ingredient, error) {
	return nil, nil
}
