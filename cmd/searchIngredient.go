package cmd

import (
	"fmt"

	"os"

	"github.com/kdelwat/recipaliser/db"
	"github.com/kdelwat/recipaliser/model"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var searchIngredientCommand = &cobra.Command{
	Use:   "search",
	Short: "Search for an ingredient",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchIngredients(args[0])
	},
}

func init() {
	ingredientCmd.AddCommand(searchIngredientCommand)
}

func searchIngredients(searchString string) {
	var matchingIngredients []model.Ingredient

	db.Db.Where("name LIKE ?", "%"+searchString+"%").Find(&matchingIngredients)

	fmt.Printf("Found %v matches\n", len(matchingIngredients))

	if len(matchingIngredients) == 0 {
		return
	}

	outputTable := tablewriter.NewWriter(os.Stdout)
	outputTable.SetHeader([]string{"ID", "Name"})

	for _, ingredient := range matchingIngredients {
		outputTable.Append([]string{fmt.Sprint(ingredient.ID), ingredient.Name})
	}

	outputTable.Render()
}
