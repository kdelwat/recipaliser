package formatters

import (
	"fmt"
	"github.com/kdelwat/recipaliser"
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintRecipe(recipe recipaliser.Recipe) {
	fmt.Printf("Name: %v\n", recipe.Name)

	outputTable := tablewriter.NewWriter(os.Stdout)

	outputTable.SetHeader([]string{"Ingredient", "Amount (g)"})

	for _, ingredient := range recipe.Ingredients {
		outputTable.Append([]string{ingredient.IngredientName, fmt.Sprintf("%v", ingredient.Amount)})
	}

	outputTable.Render()
}
