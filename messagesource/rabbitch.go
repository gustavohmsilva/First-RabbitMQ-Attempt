package messagesource

import "github.com/streadway/amqp"

// MessageChannel is the struct containing the connection to the message channel
type MessageChannel struct {
	Ch *amqp.Channel
}

// NewMessageChannel is the struct containing the channel to the message service
func NewMessageChannel(rMS MessageSource) (MessageChannel, error) {
	var rMC MessageChannel
	var err error
	rMC.Ch, err = rMS.Server.Channel()
	if err != nil {
		return MessageChannel{}, err
	}
	return rMC, nil
}
