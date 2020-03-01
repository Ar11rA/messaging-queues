package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func main() {
	conn, err1 := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err1 != nil {
		fmt.Println(err1)
		panic(true)
	}
	defer conn.Close()
	ch, err2 := conn.Channel()
	if err2 != nil {
		fmt.Println(err2)
		panic(true)
	}
	defer ch.Close()
	q, _ := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	var body = os.Args[1]
	_ = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}
