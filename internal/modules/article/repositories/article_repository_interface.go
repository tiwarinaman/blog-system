package repositories

import (
	ArticleModel "blog-system/internal/modules/article/model"
)

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Fetch(id int) ArticleModel.Article
}
