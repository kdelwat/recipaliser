package formatters

import (
	"fmt"
	"github.com/kdelwat/recipaliser"
	"github.com/olekukonko/tablewriter"
	"os"
)

type ingredientField struct {
	field string
	value float64
}

func selectIngredientFields(ingredient recipaliser.Ingredient, selections ...string) []ingredientField {
	selectionSets := map[string][]ingredientField{
		"macronutrients": {
			{field: "Energy (kJ)", value: ingredient.EnergyWithDietaryFibre},
			{field: "Available carbohydrates (with sugar alcohols)", value: ingredient.AvailableCarbohydratesWithSugarAlcohols},
			{field: "Dietary fibre", value: ingredient.DietaryFibre},
			{field: "Protein", value: ingredient.Protein},
			{field: "Total fat", value: ingredient.TotalFat},
		},
		"carbohydrates": {
			{field: "Available carbohydrates (with sugar alcohols)", value: ingredient.AvailableCarbohydratesWithSugarAlcohols},
			{field: "Available carbohydrates (without sugar alcohols)", value: ingredient.AvailableCarbohydratesWithoutSugarAlcohol},
			{field: "Starch", value: ingredient.Starch},
			{field: "Total sugars", value: ingredient.TotalSugars},
			{field: "Added sugars", value: ingredient.AddedSugars},
			{field: "Free sugars", value: ingredient.FreeSugars},
		},
		"protein": {
			{field: "Protein", value: ingredient.Protein},
			{field: "Tryptophan", value: ingredient.Tryptophan},
		},
		"fats": {
			{field: "Total fat", value: ingredient.TotalFat},
			{field: "Cholesterol", value: ingredient.Cholesterol},
			{field: "Total saturated fat", value: ingredient.TotalSaturatedFat},
			{field: "Total monounsaturated fat", value: ingredient.TotalMonounsaturatedFat},
			{field: "Total polyunsaturated fat", value: ingredient.TotalPolyunsaturatedFat},
			{field: "Linoleic acid", value: ingredient.LinoleicAcid},
			{field: "Alphalinolenic acid", value: ingredient.AlphalinolenicAcid},
			{field: "EPA", value: ingredient.C205w3Eicosapentaenoic},
			{field: "DPA", value: ingredient.C225w3Docosapentaenoic},
			{field: "DHA", value: ingredient.C226w3Docosahexaenoic},
			{field: "Total long-chain omega-3 fatty acids", value: ingredient.TotalLongChainOmega3FattyAcids},
			{field: "Total trans-fatty acids", value: ingredient.TotalTransFattyAcids},
		},
		"vitamins": {
			{field: "Vitamin A (retinol equivalents)", value: ingredient.VitaminARetinolEquivalents},
			{field: "Thiamin (B1)", value: ingredient.ThiaminB1},
			{field: "Riboflavin (B2)", value: ingredient.RiboflavinB2},
			{field: "Niacin (B3) (derived equivalents)", value: ingredient.NiacinDerivedEquivalents},
			{field: "Dietary folate equivalents", value: ingredient.DietaryFolateEquivalents},
			{field: "Vitamin B6", value: ingredient.VitaminB6},
			{field: "Vitamin B12", value: ingredient.VitaminB12},
			{field: "Vitamin C", value: ingredient.VitaminC},
			{field: "Vitamin E", value: ingredient.VitaminE},
		},
		"minerals": {
			{field: "Calcium (Ca)", value: ingredient.CalciumCa},
			{field: "Iodine (I)", value: ingredient.IodineI},
			{field: "Iron (Fe)", value: ingredient.IronFe},
			{field: "Magnesium (Mg)", value: ingredient.MagnesiumMg},
			{field: "Phosphorus (P)", value: ingredient.PhosphorusP},
			{field: "Potassium (K)", value: ingredient.PotassiumK},
			{field: "Selenium (Se)", value: ingredient.SeleniumSe},
			{field: "Sodium (Na)", value: ingredient.SodiumNa},
			{field: "Zinc (Zn)", value: ingredient.ZincZn},
		},
		"stimulants": {
			{field: "Caffeine", value: ingredient.Caffeine},
		},
		"depressants": {
			{field: "Alcohol", value: ingredient.Alcohol},
		},
	}

	var selectionSet [][]ingredientField
	for _, selection := range selections {
		selectionSet = append(selectionSet, selectionSets[selection])
	}

	return flattenSelectionSet(selectionSet)
}

func containsField(fields []ingredientField, target ingredientField) bool {
	for i := 0; i < len(fields); i++ {
		if fields[i].field == target.field {
			return true
		}
	}

	return false
}

func flattenSelectionSet(selectionSet [][]ingredientField) []ingredientField {
	var flattenedFields []ingredientField

	for i := len(selectionSet) - 1; i >= 0; i-- {
		for j := len(selectionSet[i]) - 1; j >= 0; j-- {
			if !containsField(flattenedFields, selectionSet[i][j]) {
				flattenedFields = append(flattenedFields, selectionSet[i][j])
			}
		}
	}

	// Reverse fields
	// From https://stackoverflow.com/a/42545484
	for i, j := 0, len(flattenedFields)-1; i < j; i, j = i+1, j-1 {
		flattenedFields[i], flattenedFields[j] = flattenedFields[j], flattenedFields[i]
	}

	return flattenedFields
}

func printIngredient(ingredient recipaliser.Ingredient, table *tablewriter.Table, selections ...string) {
	tableValues := []string{ingredient.Name}

	for _, f := range selectIngredientFields(ingredient, selections...) {
		tableValues = append(tableValues, fmt.Sprintf("%f", f.value))
	}

	table.Append(tableValues)
}

func PrintIngredients(ingredients []recipaliser.Ingredient, selections ...string) {
	outputTable := tablewriter.NewWriter(os.Stdout)

	tableHeaders := []string{"Name"}
	for _, f := range selectIngredientFields(ingredients[0], selections...) {
		tableHeaders = append(tableHeaders, f.field)
	}

	outputTable.SetHeader(tableHeaders)

	for _, ingredient := range ingredients {
		printIngredient(ingredient, outputTable, selections...)
	}

	outputTable.Render()
}
