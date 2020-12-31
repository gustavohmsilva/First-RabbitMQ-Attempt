package messagesource

import "github.com/streadway/amqp"

// Queue holds the data necessary to declare a new queue
type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

// NewMessageQueue declares a new queue to the channel, basically a wrapper
func NewMessageQueue(rMC MessageChannel, q Queue) (amqp.Queue, error) {
	rMQ, err := rMC.Ch.QueueDeclare(
		q.Name,
		q.Durable,
		q.AutoDelete,
		q.Exclusive,
		q.NoWait,
		q.Args,
	)
	if err != nil {
		return amqp.Queue{}, err
	}
	return rMQ, nil
}
