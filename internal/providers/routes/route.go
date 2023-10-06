package routes

import (
	articleRoutes "blog-system/internal/modules/article/routes"
	homeRoutes "blog-system/internal/modules/home/routes"
	userRoutes "blog-system/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
