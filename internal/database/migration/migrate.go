package migration

import (
	articleModels "blog-system/internal/modules/article/model"
	userModels "blog-system/internal/modules/user/models"
	"blog-system/pkg/database"
	"log"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})
	if err != nil {
		log.Fatal("Cannot migrate database tables")
		return
	}

	log.Println("Migration done")
}
