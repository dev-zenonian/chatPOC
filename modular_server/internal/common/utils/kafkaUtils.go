package utils

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func InitKafka(address string, topic string, partition int) (*kafka.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()
	log.Printf("Init Kafka connection at address: %v, topic: %v, partition: %v\n", address, topic, partition)
	conn, err := kafka.DialLeader(ctx, "tcp", address, topic, partition)
	if err != nil {
		return nil, err
	}
	log.Println("checkpoint")
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	return conn, nil
}
