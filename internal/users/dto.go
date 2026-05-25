package users

type CreateUserDTO struct {
	Name     string
	Email    string
	Password string
}

type UpdateUserDTO struct {
	Name  string
	Email string
}