package main

import (
	"encoding/json"
	"fmt"
	"log"

	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaclient "github.com/ghostriderdev/movies/pkg/messaging/kafka"
	model "github.com/ghostriderdev/movies/rating/pkg"
)

func main() {
	fmt.Println("Creating a kafka producer")

	producer, err := kafkaclient.NewProducer("localhost")

	if err != nil {
		panic(err)
	}

	const fileName = "ratingsdata.json"

	fmt.Println("Reading rating events from " + fileName)

	ratingEvents, err := readRatingEvents(fileName)

	if err != nil {
		panic(err)
	}

	const topic = "ratings"

	if err := produceRatingEvents(topic, producer, &ratingEvents); err != nil {
		panic(err)
	}

	const timeout = 10 * time.Second

	fmt.Println("Waiting " + timeout.String() + " until all events get produced")

	producer.Flush(int(timeout.Milliseconds()))
}

func produceRatingEvents(topic string, producer *kafka.Producer, ratings *[]model.RatingEvent) error {
	for _, rating := range *ratings {
		encodedEvent, err := json.Marshal(rating)

		if err != nil {
			return err
		}

		if err := kafkaclient.ProduceEvent(topic, producer, encodedEvent); err != nil {
			log.Println("Error aqui **************")
			return err
		}
	}
	return nil
}

func readRatingEvents(fileName string) ([]model.RatingEvent, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var ratings []model.RatingEvent

	if err := json.NewDecoder(f).Decode(&ratings); err != nil {
		return nil, err
	}
	return ratings, nil
}
