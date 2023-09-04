package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequest(c *gin.Context) {

	content := c.Request.Method
	if content != "GET" {
		c.JSON(http.StatusForbidden, gin.H{"message": "error customer auth", "status": "Error", "data": nil})
		c.Abort()
		return
	}
	log.Println(content)
	c.Next()
}
