package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var conn *amqp.Connection
var ch *amqp.Channel
var emailQ amqp.Queue
var err error

func Init() {
	conn, err = amqp.Dial("amqp://guest:guest@localhost:15672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	emailQ, err = ch.QueueDeclare("email", false, false, false, false, nil)
	failOnError(err, "Failed to open a queue")
}

func PublishMessageToEmailQueue(message string) {
	ctx := context.Background()
	err := ch.PublishWithContext(ctx, "", emailQ.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message)})
	failOnError(err, "Failed when send message")
	log.Printf("[x] Sent %s", message)
}

func Close() {
	defer conn.Close()
	defer ch.Close()
}
