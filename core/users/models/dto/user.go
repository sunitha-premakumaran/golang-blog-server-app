package dto

type UserCreateDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserCreateResponseDto struct {
	CreatedUserId string `json:"createdUserId"`
}
