package main

import (
	"fmt"

	"github.com/gustavohmsilva/First-RabbitMQ-Attempt/config"
)

func main() {
	fmt.Println("Go RabbitMQ - First Attempt")
	rMQ, err := config.NewMessageSource(false)
	if err != nil {
		panic(err)
	}
	defer rMQ.Server.Close()
	fmt.Println("Successfully Connected to RabbitMQ")
}
