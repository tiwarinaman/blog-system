package seeder

import (
	articleModels "blog-system/internal/modules/article/model"
	userModels "blog-system/internal/modules/user/models"
	"blog-system/pkg/database"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Seed() {
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("naman@123"), 12)
	if err != nil {
		log.Fatal("Hash password error")
		return
	}

	user := userModels.User{Name: "Naman Kumar", Email: "naman@gmail.com", Password: string(hashedPassword)}

	db.Create(&user)

	log.Printf("User created successfully with email: %s\n", user.Email)

	for i := 1; i <= 10; i++ {
		article := articleModels.Article{
			Title:   fmt.Sprintf("Randon Title %d", i),
			Content: "This the dummy content for this article",
			UserId:  1,
		}

		db.Create(&article)

		log.Printf("Article created %d\n", i)
	}

	log.Println("Seeder done...")

}
