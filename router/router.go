package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"query/constants"
	"query/controller"
	"query/message_counter"
)

func SetupRoutes() {
	httpRouter := gin.Default()

	//handling CORS
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Golang RabbitMQ API 📺 Up and Running"})
	})

	httpRouter.POST("task1", controller.Task1Controller)
	httpRouter.GET("/record-count", message_counter.GetRecordCountHandler)

	httpRouter.Run(constants.SERVERPORT)

}
