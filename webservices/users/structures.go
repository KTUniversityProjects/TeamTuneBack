package users

type RegisterStructure struct {
	Username string
	Password string
	Password2 string
	Email    string
}

type LoginStructure struct {
	Username string
	Password string
}