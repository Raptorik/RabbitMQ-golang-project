package lib

import (
	"fmt"
	"github.com/streadway/amqp"
	"query/constants"
	"query/utils"
)

var RabbitChannel *amqp.Channel
var rabbitConn *amqp.Connection

func SetupRabbbitMQConnectionChannel() (*amqp.Connection, *amqp.Channel) {

	//dial
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", constants.USERNAME, constants.PASSWORD, constants.HOST, constants.PORT)

	conn, err := amqp.Dial(url)

	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	utils.FailOnError(err, "Failed to open a channel")

	RabbitChannel = ch

	return rabbitConn, RabbitChannel

}
