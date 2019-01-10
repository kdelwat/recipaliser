package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kdelwat/recipaliser/actions"
)

func main() {
	actions.Execute()
}
