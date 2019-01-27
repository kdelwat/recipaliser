package actions

import (
	"github.com/kdelwat/recipaliser"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var addIngredientCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an ingredient to a recipe",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		initServices()

		recipeID := recipaliser.RecipeID(args[0])
		ingredientID := recipaliser.IngredientID(args[1])

		amount, err := strconv.ParseFloat(args[2], 64)

		if err != nil {
			log.Fatal("Amount must be a number")
		}

		if amount < 0 {
			log.Fatal("Amount must be positive")
		}

		if err := rs.AddIngredientToRecipe(recipeID, ingredientID, recipaliser.IngredientAmount(amount), &is); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	recipeCmd.AddCommand(addIngredientCmd)
}
