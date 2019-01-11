package actions

import (
	"fmt"

	"os"

	"github.com/kdelwat/recipaliser"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var ingredientSearchName string

var searchIngredientCommand = &cobra.Command{
	Use:   "search",
	Short: "Search for an ingredient",
	Run: func(cmd *cobra.Command, args []string) {
		initServices()
		ingredients, err := is.SearchIngredient(ingredientSearchName)

		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		printIngredients(ingredients)
	},
}

func init() {
	searchIngredientCommand.Flags().StringVarP(&ingredientSearchName, "name", "n", "", "A search term within the name of the ingredient")
	ingredientCmd.AddCommand(searchIngredientCommand)
}

func printIngredients(ingredients []recipaliser.Ingredient) {
	fmt.Printf("Found %v matches\n", len(ingredients))

	if len(ingredients) == 0 {
		return
	}

	outputTable := tablewriter.NewWriter(os.Stdout)
	outputTable.SetHeader([]string{"Name"})

	for _, i := range ingredients {
		outputTable.Append([]string{i.Name})
	}

	outputTable.Render()
}
