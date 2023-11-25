package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

func ProducerService(c echo.Context) {
	// creating producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:9092",
		"client.id":         "cli",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	// defining the name of the topic
	topic := "msg"

	deliveryChan := make(chan kafka.Event, 10000)

	// getting the data from the request body
	data := make(map[string]interface{})
	if err := c.Bind(&data); err != nil {
		log.Println(err)
	}
	fmt.Println(color.GreenString("ProducerService"))
	fmt.Println(data["message"])
	// converting the data to byte array
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	// writing message to a topic
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          dataBytes},
		deliveryChan,
	)
	if err != nil {
		log.Println(err)
	}
}
