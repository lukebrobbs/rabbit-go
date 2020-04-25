package main

import (
	"fmt"
	"os"

	"github.com/lukebrobbs/rabbit-go/cmd/stringsvc"
	"github.com/streadway/amqp"
)

// Message represents the message recieved from the queue
type Message struct {
	Message string `json:"message"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost/")

	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ", err)
		os.Exit(1)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println("Unable to create channel", err)
	}

	testQueue, err := ch.QueueDeclare(
		"test_go",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("Unable to create queue", err)
	}

	testMessages, err := ch.Consume(
		testQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("Unable to consume from queue", err)
	}
	svc := stringsvc.New()

	rabbit := stringsvc.NewAmqp(stringsvc.MakeUppercaseEndpoint(svc), stringsvc.DecodeUppercaseMessage)

	go func() {
		for m := range testMessages {
			rabbit.HandleDelivery(m)
		}
	}()

	forever := make(chan bool)

	fmt.Println("Listening")
	<-forever
}
