package api

import (
	"os"
	"github.com/gin-gonic/gin"
)

func Listen() {
	server := gin.Default()

	registerRoutes(server)

	server.Run(":" + os.Getenv("TOOLBELT_API_PORT"))
}

func registerRoutes(server *gin.Engine) {
	podRoute := server.Group("/pod")
	{
		podRoute.POST("/create", Create)
		podRoute.GET("/state/:pod", State)
	}
}
