package actions

import (
	"fmt"
	"github.com/kdelwat/recipaliser"
	"github.com/spf13/cobra"
	"log"
)

var ingredientSortField string
var ingredientSortOrder string // ascending or descending
var maxIngredients int

var listIngredientCommand = &cobra.Command{
	Use:   "list",
	Short: "List ingredients",
	Run: func(cmd *cobra.Command, args []string) {
		var sortOrder uint

		if ingredientSortOrder == "ascending" {
			sortOrder = recipaliser.Sort_Ascending
		} else if ingredientSortOrder == "descending" {
			sortOrder = recipaliser.Sort_Descending
		} else {
			log.Fatal("Invalid sort order")
		}

		initServices()

		ingredients, err := is.ListIngredients(ingredientSortField, sortOrder, maxIngredients)

		if err != nil {
			fmt.Printf("Error: %v", err)
		} else {
			printIngredients(ingredients)
		}
	},
}

func init() {
	listIngredientCommand.Flags().StringVarP(&ingredientSortField, "sortBy", "s", "name", "The name of a field to sort by")
	listIngredientCommand.Flags().StringVarP(&ingredientSortOrder, "order", "o", "descending", "The order to sort by")
	listIngredientCommand.Flags().IntVarP(&maxIngredients, "max-ingredients", "m", -1, "The max items to display")
	ingredientCmd.AddCommand(listIngredientCommand)
}
