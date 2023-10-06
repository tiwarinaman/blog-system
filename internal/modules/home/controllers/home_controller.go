package controllers

import (
	"blog-system/internal/modules/article/servies"
	"blog-system/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	articleService servies.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: servies.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":   "Home Page",
		"feature": controller.articleService.GetFeaturedArticles(),
		"stories": controller.articleService.GetStoriesArticles(),
	})
}
