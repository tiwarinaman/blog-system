package model

import (
	"blog-system/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:191"`
	Content string `gorm:"varchar:191"`
	UserId  uint
	User    models.User
}
