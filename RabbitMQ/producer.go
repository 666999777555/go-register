package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErr(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"zpy",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "Failed to declare a queue")

	body := "去你妈的"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnErr(err, "Failed to publish a message")

	log.Printf("Message sent: %s", body)
}
