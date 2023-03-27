package entities

type User struct {
	BaseModel

	FirstName string
	LastName  string
	Email     string
	Password  string
}
