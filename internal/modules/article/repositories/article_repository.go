package repositories

import (
	ArticleModel "blog-system/internal/modules/article/model"
	"blog-system/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article

	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

func (articleRepository *ArticleRepository) Fetch(id int) ArticleModel.Article {
	var article ArticleModel.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}
