package kafkaclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

// NewConsumer creates a new Kafka consumer using Sarama.
func NewConsumer(bootstrapServers, groupId, topic string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup([]string{bootstrapServers}, groupId, config)
	if err != nil {
		return nil, err
	}

	return consumerGroup, nil
}

// ConsumerGroupHandler handles Kafka messages.
type ConsumerGroupHandler[T any] struct {
	ready chan bool
	ch    chan T
}

func (h *ConsumerGroupHandler[T]) Setup(_ sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

func (h *ConsumerGroupHandler[T]) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler[T]) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var value T
		if err := json.Unmarshal(msg.Value, &value); err == nil {
			h.ch <- value
		} else {
			fmt.Printf("Unmarshal error: %v\n", err)
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}

func Consume[T any](ctx context.Context, topic string, consumerGroup sarama.ConsumerGroup) (chan T, error) {
	ch := make(chan T, 1)
	handler := ConsumerGroupHandler[T]{
		ready: make(chan bool),
		ch:    make(chan T),
	}

	go func() {
		defer close(ch)
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, &handler); err != nil {
				log.Printf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			handler.ready = make(chan bool)
		}
	}()

	go func() {
		for value := range handler.ch {
			ch <- value
		}
	}()

	<-handler.ready
	return ch, nil
}
