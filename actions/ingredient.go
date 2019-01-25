package actions

import (
	"github.com/kdelwat/recipaliser"
	"github.com/kdelwat/recipaliser/formatters"
	"github.com/spf13/cobra"
	"log"
)

var selections []string

var ingredientCmd = &cobra.Command{
	Use:   "ingredient",
	Short: "Perform operations on ingredients",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initServices()

		ingredientName := recipaliser.IngredientID(args[0])
		ingredient, err := is.Ingredient(ingredientName)

		if err != nil {
			log.Fatal(err)
		}

		formatters.PrintIngredients([]recipaliser.Ingredient{ingredient}, selections...)
	},
}

func init() {
	ingredientCmd.PersistentFlags().StringSliceVarP(&selections, "fields", "f", []string{"macronutrients", "carbohydrates", "protein", "fats", "vitamins", "minerals", "stimulants", "depressants"}, "A list of field categories to show on the ingredient")
	rootCmd.AddCommand(ingredientCmd)
}
