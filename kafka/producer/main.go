package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

func main() {
	topic := "mail"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:32772", topic, partition)

	if err != nil {
		panic(err)
	}
	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte(os.Args[1])},
	)

	conn.Close()
}
