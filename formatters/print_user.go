package formatters

import (
	"fmt"
	"github.com/kdelwat/recipaliser"
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintUser(user recipaliser.User) {
	outputTable := tablewriter.NewWriter(os.Stdout)

	outputTable.SetHeader([]string{"Field", "Value"})

	outputTable.Append([]string{"Name", user.Name})
	outputTable.Append([]string{"Age", fmt.Sprintf("%v", user.Age)})
	outputTable.Append([]string{"Pregnant?", boolToHuman(user.IsPregnant)})
	outputTable.Append([]string{"Lactating?", boolToHuman(user.IsLactating)})
	outputTable.Append([]string{"Weight", fmt.Sprintf("%v", user.Weight)})

	if user.BodyFatPercentage != -1 {
		outputTable.Append([]string{"", fmt.Sprintf("%v", user.BodyFatPercentage)})
	}

	outputTable.Render()
}

func boolToHuman(x bool) string {
	if x == true {
		return "Yes"
	} else {
		return "No"
	}
}
