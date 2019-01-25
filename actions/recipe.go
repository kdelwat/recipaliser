package actions

import (
	"github.com/kdelwat/recipaliser"
	"github.com/kdelwat/recipaliser/formatters"
	"log"

	"github.com/spf13/cobra"
)

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
	},
}

func init() {
	rootCmd.AddCommand(recipeCmd)
}
