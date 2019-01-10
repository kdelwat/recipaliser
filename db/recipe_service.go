package db

import "github.com/kdelwat/recipaliser"

var _ recipaliser.RecipeService = &RecipeService{}

type RecipeService struct {
	database *Database
}

func (rs *RecipeService) Recipe(id recipaliser.RecipeID) (*recipaliser.Recipe, error) {
	return nil, nil
}
func (rs *RecipeService) CreateRecipe(recipe *recipaliser.Recipe) error {
	return nil
}
func (rs *RecipeService) AddIngredientToRecipe(id recipaliser.RecipeID, ingredientId recipaliser.IngredientID, amount recipaliser.IngredientAmount) error {
	return nil
}
func (rs *RecipeService) RemoveIngredientFromRecipe(id recipaliser.RecipeID, ingredientId recipaliser.IngredientID) error {
	return nil
}
