package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println("Err Connect", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to rabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Err Channel", err)

	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Err Queue", err)
	}

	type Person struct {
		Name     string
		Age      int
		Email    string
		Password string
	}

	Arif := Person{
		Name:     "Arif",
		Age:      28,
		Email:    "arif@test.com",
		Password: "123456",
	}

	DataJson, _ := json.Marshal(Arif)

	if err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        DataJson,
		},
	); err != nil {
		fmt.Println("Err Publish", err)
	}

	fmt.Println("Successfully publish message to Queue")

}
