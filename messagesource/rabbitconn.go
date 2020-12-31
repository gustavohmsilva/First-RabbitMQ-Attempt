package messagesource

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

// MessageSource is the struct containing the connection to the message service
type MessageSource struct {
	Server *amqp.Connection
}

// ConnectionInfo hold the configuration for a RabbitMQ connection
type ConnectionInfo struct {
	User    string
	Pass    string
	Address string
	port    string
	Err     error
}

// NewMessageSource create a new RabbitMQ server object to deliver messages to
func NewMessageSource(prod bool) (MessageSource, error) {
	if !prod {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
		if err != nil {
			return MessageSource{}, err
		}
		var src MessageSource
		src.Server = conn
		return src, nil
	}
	errorMsgTemplate := "Missing RabbitMQ %s System Variable"
	var c ConnectionInfo
	if c.User = os.Getenv("RABBITMQ_USER"); c.User == "" {
		c.Err = fmt.Errorf(fmt.Sprintf(errorMsgTemplate, "User"))
		return MessageSource{}, c.Err
	}
	if c.User = os.Getenv("RABBITMQ_PASS"); c.User == "" {
		c.Err = fmt.Errorf(fmt.Sprintf(errorMsgTemplate, "Pass"))
		return MessageSource{}, c.Err
	}
	if c.User = os.Getenv("RABBITMQ_SERVER"); c.User == "" {
		c.Err = fmt.Errorf(fmt.Sprintf(errorMsgTemplate, "Server"))
		return MessageSource{}, c.Err
	}
	if c.User = os.Getenv("RABBITMQ_PORT"); c.User == "" {
		c.Err = fmt.Errorf(fmt.Sprintf(errorMsgTemplate, "Port"))
		return MessageSource{}, c.Err
	}
	conn, err := amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s",
			c.User,
			c.Pass,
			c.Address,
			c.port,
		),
	)
	if err != nil {
		return MessageSource{}, err
	}
	var src MessageSource
	src.Server = conn
	return src, nil
}
