package db

import (
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

type Database struct {
	Path string

	// services
	RecipeService     RecipeService
	IngredientService IngredientService

	// db connection
	session sqlbuilder.Database
}

func NewDatabase(path string) (*Database, error) {
	database := &Database{Path: path}

	database.IngredientService.database = database
	database.RecipeService.database = database

	settings := sqlite.ConnectionURL{
		Database: path,
	}

	session, err := sqlite.Open(settings)

	database.session = session

	return database, err
}

func (database *Database) Collection(name string) db.Collection {
	return database.session.Collection(name)
}

func (database *Database) Close() error {
	return nil
}
