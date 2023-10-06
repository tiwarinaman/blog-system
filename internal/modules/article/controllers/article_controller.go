package controllers

import (
	"blog-system/internal/modules/article/servies"
	"blog-system/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	articleService servies.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: servies.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{
			"title":   "Server Error",
			"message": "error converting id into int",
		})
		return
	}

	article, err := controller.articleService.Find(id)

	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{
			"title":   "Not Found",
			"message": "article not found",
		})
		return
	}

	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{
		"title":   "Show Article",
		"article": article,
	})
}
