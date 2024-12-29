package kafkaclient

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProduceEvent(topic string, producer *kafka.Producer, message []byte) error {

	log.Println("Pasa por aqui+++++++++++++++", topic)

	if err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny,
		},
		Value: []byte(message),
	}, nil); err != nil {
		return err
	}

	return nil
}

func NewProducer(bootstrapServers string) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})

	if err != nil {
		return nil, err
	}

	defer producer.Close()
	return producer, nil
}
