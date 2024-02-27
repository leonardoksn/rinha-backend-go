package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoksn/rinha-backend-go/schemas"
)

func CreateTransactionHandler(c *gin.Context) {

	var requestBody CreateTransactionRequestBody
	var requestParams CreateTransactionRequestParams

	if err := c.ShouldBindUri(&requestParams); err != nil {
		c.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	rows := db.QueryRow("SELECT id FROM clientes WHERE id = ?", requestParams.Id)

	var user schemas.User

	if err := rows.Scan(&user.Id); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := requestBody.Validate(); err != nil {
		c.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
