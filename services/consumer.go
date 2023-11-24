package services

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ConsumerService() {
	topic := "msg"
	// setting the consumer
	fmt.Println("ConsumerService")
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:9092",
		"group.id":          "go",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println("Failed to create consumer: ", err)
		return
	}
	defer consumer.Close()

	// Subscribe to the same topic as producer
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Println(err)
		return
	}

	run := true
	for run {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Printf("Received message: %s\n", string(e.Value))
		case kafka.Error:
			log.Printf("Error: %v\n", e)
			run = false
		}
	}
}
