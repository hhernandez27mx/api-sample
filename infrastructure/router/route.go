package router

import (
	"net/http"

	"azt.com/api-sample/infrastructure/controller/handler"
	"azt.com/api-sample/infrastructure/middleware"
	"azt.com/api-sample/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

func CreateRoute(p *persistence.Persistence) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//login
	contactHandler := handler.NewContactHandler(p)

	rContact := r.Group("contact")
	rContact.Use(middleware.AuthRequest)
	rContact.POST("/", handler.AddContactHandler(contactHandler))
	rContact.GET("/:id", handler.GetContactHandlerGin)
	rContact.GET("/byName", contactHandler.GetContactByNameHandler)
	return r
}
