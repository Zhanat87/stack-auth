package main

import (
	"log"
	"github.com/streadway/amqp"
	"fmt"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getConnection() (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("STACK_RABBITMQ_ENV_RABBITMQ_DEFAULT_USER"),
		os.Getenv("STACK_RABBITMQ_ENV_RABBITMQ_DEFAULT_PASS"),
		os.Getenv("STACK_RABBITMQ_PORT_15671_TCP_ADDR"),
		os.Getenv("STACK_RABBITMQ_PORT_5672_TCP_PORT")))
}

func main() {
	conn, err := getConnection()
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "hello"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
