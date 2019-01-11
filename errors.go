package recipaliser

type Error string

const (
	IngredientAlreadyExists = Error("ingredient already exists")
)

func (e Error) Error() string {
	return string(e)
}
