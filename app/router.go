package router

import (
	"net/http"
	"web/go-cat-match/controller"
	"web/go-cat-match/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.Group("/v1")
	r.POST("/user/login", controller.Login)
	r.POST("/user/register", controller.Register)

	r.POST("/cat", middleware.AuthMiddleware, controller.RegisterCat)
	r.GET("/cat", middleware.AuthMiddleware, controller.ListCat)
	r.PUT("/cat/:id", middleware.AuthMiddleware, controller.EditCat)
	r.DELETE("/cat/:id", middleware.AuthMiddleware, controller.DeleteCat)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
}
