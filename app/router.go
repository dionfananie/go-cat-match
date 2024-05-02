package router

import (
	"net/http"
	"web/go-cat-match/controller"
	"web/go-cat-match/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	router := r.Group("/v1")
	router.POST("/user/login", controller.Login)
	router.POST("/user/register", controller.Register)
	router.Use(middleware.AuthMiddleware())
	router.POST("/cat", controller.RegisterCat)
	router.GET("/cat", controller.ListCat)
	router.PUT("/cat/:id", controller.EditCat)
	router.DELETE("/cat/:id", controller.DeleteCat)
	// Route with auth middleware example (only user logged in can access this route)
	// router.GET("/v1/cat", middleware.AuthMiddleware, controller.SomeMethodHere)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
}
