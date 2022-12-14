package core

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Queue struct {
	Client  *redis.Client
	Channel *redis.PubSub
	Context context.Context
}

func NewQueue(client *redis.Client, context context.Context) *Queue {
	queuePointer := &Queue{
		Client:  client,
		Channel: nil,
		Context: context,
	}

	return queuePointer
}

func (queue Queue) OnMessageReceive(handler func(job *redis.Message)) {
	message, err := queue.Channel.ReceiveMessage(queue.Context)
	if err != nil {
		queue.Client.Publish(
			queue.Context,
			"teste",
			"teste",
		)
	}
	handler(message)
}
