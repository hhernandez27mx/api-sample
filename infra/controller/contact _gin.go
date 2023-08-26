package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"azt.com/api-sample/model/entity"
	"github.com/gin-gonic/gin"
)

func AdddContactHandlerGin(c *gin.Context) {
	var data entity.Contact
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Procesar los datos recibidos
	result := fmt.Sprintf("Hola, %s. Tienes %d años y tu correo es %s.", data.Name, data.Age, data.Email)

	response := ResponseApi{0, result}
	c.JSON(http.StatusOK, response)
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
