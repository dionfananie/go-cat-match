package router

import (
	"net/http"
	"web/go-cat-match/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/v1/user/login", controller.Login)
	r.POST("/v1/user/register", controller.Register)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
}
