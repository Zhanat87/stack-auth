package rabbitmq

import (
	"log"
	"github.com/streadway/amqp"
	"fmt"
	"os"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetConnection() (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("STACK_RABBITMQ_ENV_RABBITMQ_DEFAULT_USER"),
		os.Getenv("STACK_RABBITMQ_ENV_RABBITMQ_DEFAULT_PASS"),
		os.Getenv("STACK_RABBITMQ_PORT_15671_TCP_ADDR"),
		os.Getenv("STACK_RABBITMQ_PORT_5672_TCP_PORT")))
}