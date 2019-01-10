package actions

import (
	"log"

	"github.com/kdelwat/recipaliser"
	"github.com/spf13/cobra"
)

var newRecipeCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new recipe",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := rs.CreateRecipe(&recipaliser.Recipe{Name: args[0]}); err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	recipeCmd.AddCommand(newRecipeCmd)
}
