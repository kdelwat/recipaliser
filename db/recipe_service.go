package db

import (
	"github.com/kdelwat/recipaliser"
	"log"
	"upper.io/db.v3"
)

var _ recipaliser.RecipeService = &RecipeService{}

type RecipeService struct {
	database *Database
}

func (rs *RecipeService) Recipe(id recipaliser.RecipeID) (*recipaliser.Recipe, error) {
	return nil, nil
}
func (rs *RecipeService) CreateRecipe(recipe *recipaliser.Recipe) error {
	var existingRecipe recipaliser.Recipe

	if err := rs.database.Collection("recipes").Find("name = ?", recipe.Name).One(&existingRecipe); err != nil {
		if err.Error() != "upper: no more rows in this result set" {
			return err
		}
	} else {
		return recipaliser.RecipeAlreadyExists
	}
	_, err := rs.database.Collection("recipes").Insert(*recipe)

	return err
}

func (rs *RecipeService) AddIngredientToRecipe(id recipaliser.RecipeID, ingredientId recipaliser.IngredientID, amount recipaliser.IngredientAmount) error {
	recipeIngredient := recipaliser.RecipeIngredient{RecipeName: string(id), IngredientName: string(ingredientId), Amount: float64(amount)}

	existingRecipeIngredients := rs.database.Collection("recipe_ingredients").Find(db.And(db.Cond{"recipe_name": id}, db.Cond{"ingredient_name": ingredientId}))

	existingCount, err := existingRecipeIngredients.Count()

	if err != nil {
		log.Fatal(err)
	}

	if existingCount == 0 {
		_, err := rs.database.Collection("recipe_ingredients").Insert(recipeIngredient)

		return err
	} else {
		err := existingRecipeIngredients.Update(recipeIngredient)

		return err
	}
}

func (rs *RecipeService) RemoveIngredientFromRecipe(id recipaliser.RecipeID, ingredientId recipaliser.IngredientID) error {
	return nil
}
