package responses

import (
	ArticleModel "blog-system/internal/modules/article/model"
	UserResponse "blog-system/internal/modules/user/responses"
	"fmt"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      UserResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article ArticleModel.Article) Article {
	return Article{
		ID:      article.ID,
		Image:   "/assets/img/demopic/1.jpg",
		Title:   article.Title,
		Content: article.Content,
		CreatedAt: fmt.Sprintf(
			"%d/%02d/%02d",
			article.CreatedAt.Year(),
			article.CreatedAt.Month(),
			article.CreatedAt.Day(),
		),
		User: UserResponse.ToUser(article.User),
	}
}

func ToArticles(articles []ArticleModel.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
