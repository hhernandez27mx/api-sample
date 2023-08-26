package router

import (
	"net/http"

	"azt.com/api-sample/infra/controller"
	"azt.com/api-sample/infra/middleware"
	"github.com/gin-gonic/gin"
)

func CreateRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//login

	rContact := r.Group("contact")
	rContact.Use(middleware.AuthRequest)
	rContact.POST("/", controller.AdddContactHandlerGin)
	rContact.GET("/:id", controller.GetContactHandlerGin)
	return r
}
