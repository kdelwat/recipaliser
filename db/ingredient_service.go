package db

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/kdelwat/recipaliser"
)

var _ recipaliser.IngredientService = &IngredientService{}

type IngredientService struct {
	database *Database
}

func (is *IngredientService) Ingredient(id recipaliser.IngredientID) (recipaliser.Ingredient, error) {
	var ingredient recipaliser.Ingredient

	if err := is.database.Collection("ingredients").Find("name = ?", id).One(&ingredient); err != nil {
		if err.Error() == "upper: no more rows in this result set" {
			return recipaliser.Ingredient{}, recipaliser.IngredientNotFound
		} else {
			return recipaliser.Ingredient{}, err
		}
	}

	return ingredient, nil
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

func (is *IngredientService) SearchIngredient(nameSubstring string) ([]recipaliser.Ingredient, error) {
	var ingredients []recipaliser.Ingredient

	err := is.database.Collection("ingredients").Find("name LIKE ?", "%"+nameSubstring+"%").All(&ingredients)

	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (is *IngredientService) ListIngredients(sortField string, sortOrder uint, maxIngredients int) ([]recipaliser.Ingredient, error) {
	var ingredients []recipaliser.Ingredient

	var orderBy string
	if sortOrder == recipaliser.Sort_Ascending {
		orderBy = fmt.Sprintf("%v ASC", strcase.ToSnake(sortField))
	} else {
		orderBy = fmt.Sprintf("%v DESC", strcase.ToSnake(sortField))
	}

	var err error

	if maxIngredients != -1 {
		err = is.database.Collection("ingredients").Find().OrderBy(orderBy).Limit(maxIngredients).All(&ingredients)

	} else {
		err = is.database.Collection("ingredients").Find().OrderBy(orderBy).All(&ingredients)
	}

	if err != nil {
		return nil, err
	}

	return ingredients, nil

}
