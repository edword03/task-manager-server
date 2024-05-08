package user

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"task-manager/internal/domain/entities"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) Create(user *entities.User) (*entities.User, error) {
	dbUser := ToDBUser(user)
	if err := u.db.Create(dbUser).Error; err != nil {
		return nil, err
	}

	return ToDomainUser(dbUser), nil
}

func (u UserRepository) FindById(id string) (*entities.User, error) {
	var user *User
	err := u.db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return nil, err
	}

	return ToDomainUser(user), nil
}

func (u UserRepository) ComparePassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u UserRepository) FindByUsername(username string) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user *User

	err := u.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		logrus.Error("user repo: ", err)
		return nil, err
	}

	if user.Email == "" {
		return nil, nil
	}

	return ToDomainUser(user), nil
}

func (u UserRepository) FindAll(query string) ([]*entities.User, error) {
	var users []*User
	err := u.db.Where("first_name = ", query).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return ToDomainUsers(users), nil
}

func (u UserRepository) Update(user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
