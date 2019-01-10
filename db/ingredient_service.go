package db

import "github.com/kdelwat/recipaliser"

var _ recipaliser.IngredientService = &IngredientService{}

type IngredientService struct {
	database *Database
}

func (is *IngredientService) Ingredient(id recipaliser.IngredientID) (*recipaliser.Ingredient, error) {
	return nil, nil
}

func (is *IngredientService) CreateIngredient(ingredient *recipaliser.Ingredient) error {
	return nil
}

func (is *IngredientService) SearchIngredient(nameSubstring string) ([]*recipaliser.Ingredient, error) {
	return nil, nil
}
