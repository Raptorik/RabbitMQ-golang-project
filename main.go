package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"query/constants"
	"query/lib"
	"query/message_counter"
	"query/router"
	"query/utils"
)

func main() {
	connection, channel := lib.SetupRabbbitMQConnectionChannel()
	defer connection.Close()
	defer channel.Close()

	requestQueue, err := channel.QueueDeclare(
		constants.QUEUE, // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)

	utils.FailOnError(err, "Failed to register a queue")

	request, err := channel.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)

	utils.FailOnError(err, "Failed to register a listener in queue")

	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disabled")
	if err != nil {
		utils.FailOnError(err, "Failed to connect to database")
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS messages (id serial primary key, message_type text, time_received timestamp DEFAULT current_timestamp)")
	if err != nil {
		utils.FailOnError(err, "Failed to create messages table")
	}

	go func() {
		for d := range request {
			message_counter.IncrementCounter()

			_, err = db.Exec("INSERT INTO messages (message_type) VALUES ($1)", d.Type)
			if err != nil {
				utils.FailOnError(err, "Failed to insert message into database")
			}

			switch d.Type {
			case constants.TASK1:
				fmt.Println("In progress - TASK1")
				d.Ack(false)
			}
		}

	}()

	router.SetupRoutes()

}
