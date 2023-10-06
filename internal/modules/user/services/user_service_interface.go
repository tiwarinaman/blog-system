package services

import (
	"blog-system/internal/modules/user/requests/auth"
	UserResponse "blog-system/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
}
