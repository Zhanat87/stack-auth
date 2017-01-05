package main

import (
	"log"
	"github.com/streadway/amqp"
	"fmt"
	"os"
)

func failOnError2(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getConnection2() (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("STACK_RABBITMQ_ENV_RABBITMQ_DEFAULT_USER"),
		os.Getenv("STACK_RABBITMQ_ENV_RABBITMQ_DEFAULT_PASS"),
		os.Getenv("STACK_RABBITMQ_PORT_15671_TCP_ADDR"),
		os.Getenv("STACK_RABBITMQ_PORT_5672_TCP_PORT")))
}

func main() {
	conn, err := getConnection2()
	failOnError2(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError2(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError2(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError2(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Print(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}