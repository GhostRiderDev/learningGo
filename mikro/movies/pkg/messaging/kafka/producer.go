package kafkaclient

import (
	"log"

	"github.com/IBM/sarama"
)

func ProduceEvent(topic string, producer sarama.SyncProducer, message []byte) error {

	log.Println("Pasa por aqui+++++++++++++++", topic)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	if _, _, err := producer.SendMessage(msg); err != nil {
		return err
	}

	return nil
}

func NewProducer(bootstrapServers string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{bootstrapServers}, config)

	if err != nil {
		return nil, err
	}

	return producer, nil
}
