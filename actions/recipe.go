package actions

import (
	"github.com/kdelwat/recipaliser"
	"github.com/kdelwat/recipaliser/formatters"
	"log"

	"github.com/spf13/cobra"
)

var recipeNutritionSelections []string
var showNutrition bool

var recipeCmd = &cobra.Command{
	Use:   "recipe",
	Short: "Perform operations on recipes",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initServices()

		recipeName := recipaliser.RecipeID(args[0])
		recipe, err := rs.Recipe(recipeName)

		if err != nil {
			log.Fatal(err)
		}

		formatters.PrintRecipe(recipe)

		if showNutrition {
			recipeIngredients, err := rs.RecipeIngredients(recipeName, &is)

			if err != nil {
				log.Fatal(err)
			}

			nutrition, err := recipaliser.CalculateNutritionalData(recipe, recipeIngredients)

			if err != nil {
				log.Fatal(err)
			}

			formatters.PrintRecipeNutrition(nutrition, recipeNutritionSelections...)
		}
	},
}

func init() {
	recipeCmd.Flags().StringSliceVarP(&recipeNutritionSelections, "fields", "f", []string{"macronutrients", "carbohydrates", "protein", "fats", "vitamins", "minerals", "stimulants", "depressants"}, "A list of field categories to show on the recipe")
	recipeCmd.Flags().BoolVarP(&showNutrition, "nutrition", "n", false, "Print nutritional information for a recipe")
	rootCmd.AddCommand(recipeCmd)
}
