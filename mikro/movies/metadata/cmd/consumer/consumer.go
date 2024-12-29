package main

import (
	"context"
	"log"

	kafkaclient "github.com/ghostriderdev/movies/pkg/messaging/kafka"
	model "github.com/ghostriderdev/movies/rating/pkg"
)

func main() {
	log.Println("Starting to reading messages")

	const topic = "ratings"

	consumer, err := kafkaclient.NewConsumer("localhost", "metadata", topic)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chanel, err := kafkaclient.Consume[model.RatingEvent](ctx, topic, consumer)

	if err != nil {
		log.Println("error to receive messages: ", err.Error())
	}

	for e := range chanel {
		log.Println("Value", e.Value)
		log.Println("RatingType", e.EventType)
		log.Println("UserId", e.UserID)
		log.Println("------------------------")
	}

}
