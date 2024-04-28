package dto

type LoginDTO struct {
	Email    string
	Password string
}

type RegisterDTO struct {
	Email     string
	Username  string
	FirstName string
	LastName  string
	Password  string
	Sphere    string
	Avatar    string
}
