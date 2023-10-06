package repositories

import userModels "blog-system/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user userModels.User) userModels.User
}
