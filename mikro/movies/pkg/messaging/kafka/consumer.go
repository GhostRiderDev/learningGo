package kafkaclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewConsumer creates a new Kafka consumer.
func NewConsumer(bootstrapServers, groupId, topic string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func Consume[T any](ctx context.Context, topic string, c *kafka.Consumer) (chan T, error) {
	if err := c.SubscribeTopics([]string{topic}, nil); err != nil {
		return nil, err
	}

	ch := make(chan T, 1)

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := c.ReadMessage(-1)
				if err == nil {
					var value T
					if err := json.Unmarshal(msg.Value, &value); err == nil {
						ch <- value
					} else {
						// Handle unmarshal error
						fmt.Printf("Unmarshal error: %v\n", err)
					}
				} else {
					// Handle read message error
					fmt.Printf("Consumer error: %v (%v)\n", err, msg)
				}
			}
		}
	}()

	return ch, nil
}
