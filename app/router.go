package router

import (
	"net/http"
	"web/go-cat-match/controller"
	"web/go-cat-match/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/v1/user/login", controller.Login)
	r.POST("/v1/user/register", controller.Register)

	r.POST("/v1/cat", middleware.AuthMiddleware, controller.RegisterCat)
	r.GET("/v1/cat", controller.ListCat)
	r.PUT("/v1/cat/:id", controller.EditCat)
	r.DELETE("/v1/cat/:id", middleware.AuthMiddleware, controller.DeleteCat)
	// Route with auth middleware example (only user logged in can access this route)
	// r.GET("/v1/cat", middleware.AuthMiddleware, controller.SomeMethodHere)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
}
