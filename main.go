package main

import (
	"fmt"

	"github.com/gustavohmsilva/First-RabbitMQ-Attempt/messagesource"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ - First Attempt")
	rConn, err := messagesource.NewMessageSource(false)
	if err != nil {
		panic(err)
	}
	defer rConn.Server.Close()
	rChann, err := messagesource.NewMessageChannel(rConn)
	if err != nil {
		panic(err)
	}
	defer rChann.Ch.Close()
	mq := messagesource.Queue{"Hello Test", false, false, false, false, nil}
	rQueue, err := messagesource.NewMessageQueue(rChann, mq)
	if err != nil {
		panic(err)
	}
	messageBody := "Ola, hello, Hola, Nihao, ohayou"
	err = rChann.Ch.Publish(
		"",
		rQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageBody),
		},
	)
	fmt.Println("Successfully Connected to RabbitMQ")
}
