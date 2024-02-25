package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	initializeRoutes(r)
	r.Run(":8080")
	return r
}