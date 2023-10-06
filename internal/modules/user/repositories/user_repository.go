package repositories

import (
	userModels "blog-system/internal/modules/user/models"
	"blog-system/pkg/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (userRepository *UserRepository) Create(user userModels.User) userModels.User {
	var newUser userModels.User

	userRepository.DB.Create(&user).Scan(&newUser)

	return newUser
}
