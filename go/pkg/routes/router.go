package routes

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	router := gin.Default()
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/token", Token)
	return router
}
