package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoksn/rinha-backend-go/handler"
)

func initializeRoutes(r *gin.Engine) {

	r.POST("/clientes/:id/transacoes", handler.CreateTransactionHandler)
	r.GET("/clientes/:id/extrato", handler.ConsultTransactionHandler)	
}
