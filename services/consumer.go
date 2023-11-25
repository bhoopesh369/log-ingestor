package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bhoopesh369/log-injestor/config"
	"github.com/bhoopesh369/log-injestor/models"
	"github.com/bhoopesh369/log-injestor/utils"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/labstack/echo/v4"
)

func ConsumerService(c echo.Context) error {
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
		return err
	}
	defer consumer.Close()

	// Subscribe to the same topic as producer
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	for {
		// read messages from the topic
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			// return c.String(200, string(msg.Value))
			log := new(models.Log)
			err = json.Unmarshal(msg.Value, &log)

			if err != nil {
				return utils.SendResponse(c, http.StatusInternalServerError, "Error while unmarshalling log")
			}

			db := config.GetDB()

			logCollection := db.Collection(models.LogCollectionName())
			fmt.Println(log)
			_, err = logCollection.InsertOne(c.Request().Context(), log)

			if err != nil {
				return utils.SendResponse(c, http.StatusInternalServerError, "Error while inserting log")
			}
			return utils.SendResponse(c, http.StatusOK, "Log inserted successfully")
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	consumer.Close()
	return nil
}
