package services

import (
	userModels "blog-system/internal/modules/user/models"
	"blog-system/internal/modules/user/repositories"
	"blog-system/internal/modules/user/requests/auth"
	UserResponse "blog-system/internal/modules/user/responses"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: repositories.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModels.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hashing password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating the user")
	}

	return UserResponse.ToUser(newUser), nil
}
