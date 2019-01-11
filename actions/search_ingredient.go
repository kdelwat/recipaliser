package actions

import (
	"fmt"

	"os"

	"github.com/kdelwat/recipaliser"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var searchIngredientCommand = &cobra.Command{
	Use:   "search",
	Short: "Search for an ingredient",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ingredients, err := is.SearchIngredient(args[0])

		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		printIngredients(ingredients)
	},
}

func init() {
	ingredientCmd.AddCommand(searchIngredientCommand)
}

func printIngredients(ingredients []*recipaliser.Ingredient) {
	fmt.Printf("Found %v matches\n", len(ingredients))

	if len(ingredients) == 0 {
		return
	}

	outputTable := tablewriter.NewWriter(os.Stdout)
	outputTable.SetHeader([]string{"ID", "Name"})

	for _, i := range ingredients {
		outputTable.Append([]string{i.Name})
	}

	outputTable.Render()
}
