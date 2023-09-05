package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"azt.com/api-sample/infrastructure/implementations/services"
	"azt.com/api-sample/infrastructure/persistence"
	"azt.com/api-sample/model/entity"
	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	Persistence *persistence.Persistence
}

// constructor
func NewContactHandler(p *persistence.Persistence) *ContactHandler {
	return &ContactHandler{p}
}

func (p *ContactHandler) AdddContactHandlerGin(c *gin.Context) {
	var data entity.Contact
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contctServ := services.NewContactService(p.Persistence)
	err = contctServ.AddServ(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0})
}

func GetContactHandlerGin(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Procesar los datos recibidos
	result := fmt.Sprintf("Hola, id %d.", id)

	response := ResponseApi{0, result}
	c.JSON(http.StatusOK, response)
}

func (p *ContactHandler) GetContactByNameHandler(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name empty"})
		return
	}
	contctServ := services.NewContactService(p.Persistence)
	// Procesar los datos recibidos
	result := contctServ.GetByName(name)

	response := gin.H{"data": result}
	c.JSON(http.StatusOK, response)
}

func AddContactHandler(handler *ContactHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.AdddContactHandlerGin(c)
	}
}
