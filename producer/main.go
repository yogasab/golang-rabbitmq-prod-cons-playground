package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Starting to publish message to queue ...")

	// Channel consist of producer/consumer, for this file it is producer
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		message := amqp091.Publishing{
			Headers: amqp091.Table{
				"sample": "value",
			},
			Body: []byte("Email " + strconv.Itoa(i)),
		}
		if err := channel.PublishWithContext(
			context.Background(), // context
			"notification",       // exchange
			"email",              // routing key
			false,
			false,
			message,
		); err != nil {
			panic(err)
		}
	}

	fmt.Println("Complete publish message to queue ...")
}
