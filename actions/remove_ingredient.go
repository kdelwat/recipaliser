package actions

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeIngredientCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an ingredient from a recipe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	recipeCmd.AddCommand(removeIngredientCmd)
}
