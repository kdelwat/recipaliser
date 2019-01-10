package actions

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyseRecipeCommand = &cobra.Command{
	Use:   "analyse",
	Short: "Analyse the nutrition of an existing recipe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analyse called")
	},
}

func init() {
	recipeCmd.AddCommand(analyseRecipeCommand)
}
