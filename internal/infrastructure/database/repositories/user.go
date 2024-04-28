package repositories

import (
	"gorm.io/gorm"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/repositories"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) Create(user *entities.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}

	return nil
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
	//TODO implement me
	panic("implement me")
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
