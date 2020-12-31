package main

import (
	"fmt"

	"github.com/gustavohmsilva/First-RabbitMQ-Attempt/config"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ - First Attempt")
	conf := config.NewRabbitMQ(false)
	if conf.Err != nil {
		panic(conf.Err)
	}
	connString, err := conf.ParseConnectionString()
	if err != nil {
		panic(err)
	}
	var conn *amqp.Connection
	conn, err = amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Successfully Connected to RabbitMQ")
}
