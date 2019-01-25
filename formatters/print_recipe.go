package formatters

import (
	"fmt"
	"github.com/kdelwat/recipaliser"
)

func PrintRecipe(recipe recipaliser.Recipe) {
	fmt.Printf("Name: %v\n", recipe.Name)
}
