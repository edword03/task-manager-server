package entities

type User struct {
	ID        [16]byte
	Email     string
	Username  string
	FirstName string
	LastName  string
	Password  string
	Sphere    string
	Avatar    string
}

func New(id [16]byte, username, surname, email, password string) *User {
	return &User{
		ID:       id,
		Username: username,
		LastName: surname,
		Email:    email,
		Password: password,
	}
}
