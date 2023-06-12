package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"kafka-stream-golang/config"
	"kafka-stream-golang/models"
	"log"
	"time"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.Server, config.Topic, config.Partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	msg := models.Order{
		Item:        "foo",
		Quantity:    1,
		DeliverType: "biking",
	}

	msgByte, _ := json.Marshal(msg)

	_, err = conn.WriteMessages(
		kafka.Message{Value: msgByte},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
