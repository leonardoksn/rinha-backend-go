package handler

import "github.com/gin-gonic/gin"

func ConsultTransactionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
