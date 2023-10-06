package servies

import (
	ArticleResponses "blog-system/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponses.Articles
	GetStoriesArticles() ArticleResponses.Articles
	Find(id int) (ArticleResponses.Article, error)
}
