package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:32772"},
		Topic:     "mail",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(0)
	defer r.Close()

	forever := make(chan bool)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
	<-forever
}
