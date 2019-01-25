package actions

import (
	"github.com/kdelwat/recipaliser"
	"github.com/spf13/cobra"
	"log"
)

var removeIngredientCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an ingredient from a recipe",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		initServices()

		recipeID := recipaliser.RecipeID(args[0])
		ingredientID := recipaliser.IngredientID(args[1])

		if err := rs.RemoveIngredientFromRecipe(recipeID, ingredientID); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	recipeCmd.AddCommand(removeIngredientCmd)
}
