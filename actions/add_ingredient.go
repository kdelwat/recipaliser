package actions

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addIngredientCmd represents the add command
var addIngredientCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an ingredient to a recipe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	recipeCmd.AddCommand(addIngredientCmd)
}
