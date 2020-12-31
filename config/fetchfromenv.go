package config

import (
	"fmt"
	"os"
)

// RabbitMQ hold the configuration for a RabbitMQ connection
type RabbitMQ struct {
	User   string
	Pass   string
	server string
	port   string
	Err    error
}

// NewRabbitMQ fetch the RabbitMQ configuration from system variables or create
// a dummy configuration for local environment.
func NewRabbitMQ(testing bool) RabbitMQ {
	if !testing {
		return RabbitMQ{"guest", "guest", "localhost", "5672", nil}
	}
	var c RabbitMQ
	errorMsgTemplate := "Missing RabbitMQ %s System Variable"
	if c.User = os.Getenv("RABBITMQ_USER"); c.User == "" {
		c.Err = fmt.Errorf(fmt.Sprintf(errorMsgTemplate, "User"))
		return c
	}
	if c.User = os.Getenv("RABBITMQ_PASS"); c.User == "" {
		c.Err = fmt.Errorf(fmt.Sprintf(errorMsgTemplate, "Pass"))
		return c
	}
	if c.User = os.Getenv("RABBITMQ_SERVER"); c.User == "" {
		c.Err = fmt.Errorf(fmt.Sprintf(errorMsgTemplate, "Server"))
		return c
	}
	if c.User = os.Getenv("RABBITMQ_PORT"); c.User == "" {
		c.Err = fmt.Errorf(fmt.Sprintf(errorMsgTemplate, "Port"))
		return c
	}
	return c
}

// ParseConnectionString create a string used in amqp.Dial from a RabbitMQ
// configuration struct
func (c *RabbitMQ) ParseConnectionString() (string, error) {
	if c.Err != nil {
		return "", c.Err
	}
	connectionString := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		c.User,
		c.Pass,
		c.server,
		c.port,
	)
	return connectionString, nil
}
