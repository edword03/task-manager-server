package repositories

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/repositories"
	"task-manager/internal/infrastructure/database/postgres/mappers"
	"task-manager/internal/infrastructure/database/postgres/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) Create(user *entities.User) (*entities.User, error) {
	dbUser := mappers.ToDBUser(user)
	if err := u.db.Create(dbUser).Error; err != nil {
		return nil, err
	}

	return mappers.ToDomainUser(dbUser), nil
}

func (u UserRepository) FindById(id string) (*entities.User, error) {
	var user *model.User
	err := u.db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return nil, err
	}

	return mappers.ToDomainUser(user), nil
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
	var user *model.User

	err := u.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		logrus.Error("user repo: ", err)
		return nil, err
	}

	if user.Email == "" {
		return nil, nil
	}

	return mappers.ToDomainUser(user), nil
}

func (u UserRepository) FindAll(query string) ([]*entities.User, error) {
	var users []*model.User
	err := u.db.Where("first_name = ", query).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return mappers.ToDomainUsers(users), nil
}

func (u UserRepository) Update(user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
