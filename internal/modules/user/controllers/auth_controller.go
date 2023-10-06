package controllers

import (
	"blog-system/internal/modules/user/requests/auth"
	"blog-system/internal/modules/user/services"
	"blog-system/pkg/errors"
	"blog-system/pkg/html"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct {
	userService services.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: services.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {

	var request auth.RegisterRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromError(err)

		c.Redirect(http.StatusFound, "/register")
		log.Print("Error : ", errors.Get())
		return
	}

	newUser, err := controller.userService.Create(request)

	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusFound, "/register")
	}

	log.Printf("User created successfully with name %s \n", newUser.Name)
	c.Redirect(http.StatusFound, "/")
}
