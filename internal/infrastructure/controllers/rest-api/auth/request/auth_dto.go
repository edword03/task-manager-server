package request

type RegisterDTO struct {
	Email     string `validate:"required,email"`
	Username  string `validate:"required,alphanum"`
	FirstName string `validate:"omitempty,alpha"`
	LastName  string `validate:"omitempty,alpha"`
	Password  string `validate:"required,min=8"`
	Sphere    string `validate:"required"`
	Avatar    string
}

type LoginDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
