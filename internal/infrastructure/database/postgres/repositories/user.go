package repositories

import (
	"github.com/sirupsen/logrus"
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
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) ComparePassword(password, passwordDto string) (bool, error) {
	//TODO implement me
	panic("implement me")
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

func (u UserRepository) FindAll() ([]*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Update(user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
