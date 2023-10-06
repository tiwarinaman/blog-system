package servies

import (
	ArticleRepository "blog-system/internal/modules/article/repositories"
	ArticleResponses "blog-system/internal/modules/article/responses"
	"errors"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() ArticleResponses.Articles {
	articles := articleService.articleRepository.List(4)
	return ArticleResponses.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() ArticleResponses.Articles {
	articles := articleService.articleRepository.List(6)
	return ArticleResponses.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (ArticleResponses.Article, error) {
	var response ArticleResponses.Article

	article := articleService.articleRepository.Fetch(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return ArticleResponses.ToArticle(article), nil
}
