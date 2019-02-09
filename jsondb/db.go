package jsondb

type Database struct {
	Path string

	UserService UserService
}

func NewDatabase(path string) (*Database, error) {
	database := &Database{Path: path}

	database.UserService.database = database

	return database, nil
}
