package router

import (
	"net/http"
	"web/go-cat-match/controller"
	"web/go-cat-match/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router := r.Group("/v1")
	router.POST("/user/login", controller.Login)
	router.POST("/user/register", controller.Register)

	router.Use(middleware.AuthMiddleware())
	router.POST("/cat", controller.RegisterCat)
	router.GET("/cat", controller.ListCat)
	router.PUT("/cat/:id", controller.EditCat)
	router.DELETE("/cat/:id", controller.DeleteCat)

	router.POST("/cat/match/approve", controller.ApproveMatch)
	router.POST("/cat/match/reject", controller.RejectMatch)
	router.DELETE("/cat/match/:id", controller.DeleteMatch)
	router.POST("/cat/match", controller.MatchCat)
	router.GET("/cat/match", controller.ListMatch)

}
