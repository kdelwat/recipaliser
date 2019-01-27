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

func (rs *RecipeService) Recipe(id recipaliser.RecipeID) (recipaliser.Recipe, error) {
	var recipe recipaliser.Recipe

	if err := rs.database.Collection("recipes").Find("name = ?", id).One(&recipe); err != nil {
		if err.Error() == "upper: no more rows in this result set" {
			return recipaliser.Recipe{}, recipaliser.RecipeNotFound
		} else {
			return recipaliser.Recipe{}, err
		}
	}

	if err := rs.database.Collection("recipe_ingredients").Find("recipe_name = ?", id).All(&recipe.Ingredients); err != nil {
		return recipaliser.Recipe{}, err
	}

	return recipe, nil
}

func (rs *RecipeService) RecipeIngredients(id recipaliser.RecipeID, is *recipaliser.IngredientService) ([]recipaliser.Ingredient, error) {
	var recipeIngredients []recipaliser.RecipeIngredient

	if err := rs.database.Collection("recipe_ingredients").Find("recipe_name = ?", id).All(&recipeIngredients); err != nil {
		return []recipaliser.Ingredient{}, err
	}

	var ingredients []recipaliser.Ingredient
	for _, recipeIngredient := range recipeIngredients {

		ingredient, err := (*is).Ingredient(recipaliser.IngredientID(recipeIngredient.IngredientName))

		if err != nil {
			return []recipaliser.Ingredient{}, err
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil

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

func (rs *RecipeService) AddIngredientToRecipe(id recipaliser.RecipeID, ingredientId recipaliser.IngredientID, amount recipaliser.IngredientAmount, is *recipaliser.IngredientService) error {
	// Check that ingredient doesn't exist
	if _, err := (*is).Ingredient(ingredientId); err != nil {
		return err
	}

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
	existingRecipeIngredients := rs.database.Collection("recipe_ingredients").Find(db.And(db.Cond{"recipe_name": id}, db.Cond{"ingredient_name": ingredientId}))

	return existingRecipeIngredients.Delete()
}
