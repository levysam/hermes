package main

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/levysam/hermes/core"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	context := context.Background()
	queue := core.NewQueue(client, context)
	queue.Channel = queue.Client.PSubscribe(context)

	queue.OnMessageReceive(handler)
}

func handler(message *redis.Message) {
	//do something
}
