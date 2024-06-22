package user

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"task-manager/internal/database/postgres/models"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) Create(user *entities.User) (*entities.User, error) {
	dbUser := ToDBUser(user)
	if err := u.db.Model(&models.User{}).Create(dbUser).Error; err != nil {
		return nil, err
	}

	return ToDomainUser(dbUser), nil
}

func (u UserRepository) FindById(id string) (*entities.User, error) {
	var user *models.User
	err := u.db.Model(&models.User{}).Where("id = ?", id).Find(&user).Error

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
	var user *models.User
	result := u.db.Model(models.User{}).Where("name = ?", username).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return ToDomainUser(user), nil
}

func (u UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user *models.User

	err := u.db.Model(&models.User{}).Where("email = ?", email).Find(&user).Error

	if err != nil {
		logrus.Error("user repo: ", err)
		return nil, err
	}

	return ToDomainUser(user), nil
}

func (u UserRepository) FindAll(page, pageSize int, searchTerm string) ([]*entities.User, error) {
	offset := (page - 1) * pageSize
	var users []*models.User

	query := u.db.Offset(offset).Limit(pageSize)
	if searchTerm != "" {
		query = query.Where("first_name LIKE ?", "%"+searchTerm+"%")
	}

	err := query.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return ToDomainUsers(users), nil
}

func (u UserRepository) Update(userId string, user *dto.UserDTO) error {
	result := u.db.Model(&models.User{}).Where("id = ?", userId).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println(result)

	return nil
}

func (u UserRepository) Delete(id string) error {
	result := u.db.Model(&models.User{}).Delete(&models.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
