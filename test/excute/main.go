package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://chen:0821@localhost:8080/")
	boo := failOnError(err, "Failed to connect to RabbitMQ")
	if boo {
		return
	}
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
	boo = failOnError(err, "Failed to declare a queue")
	if boo {
		return
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	boo = failOnError(err, "Failed to register a consumer")
	if boo {
		return
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) bool{
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		return true
	}
	return false
}
