package main

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Starting to consume from queue ...")

	// Channel consist of producer/consumer, for this file it is consumer
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	// Consume in RabbitMQ using Round-Robin algorithm
	emailConsumers, err := channel.ConsumeWithContext(
		context.Background(), // context
		"email",              // queue
		"consumer-email",     // consumer,
		true,                 // auto-ack set to true
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for message := range emailConsumers {
		fmt.Println("Routing Key : ", message.RoutingKey)
		fmt.Println("Body : ", string(message.Body))
		fmt.Println("Headers : ", message.Headers)
	}
}
