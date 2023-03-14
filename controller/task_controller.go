package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"net/http"
	"query/constants"
	"query/lib"
	"query/utils"
)

func Task1Controller(c *gin.Context) {
	ch := lib.RabbitChannel
	for i := 0; i < 1000000; i++ {
		err := ch.Publish(
			"",              // exchange
			constants.QUEUE, // routing key
			false,           // mandatory
			false,           // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        nil,
				Type:        constants.TASK1,
			})

		utils.FailOnError(err, "Failed to publish a message")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task-1: 1 million messages sent successfully",
	})
}
