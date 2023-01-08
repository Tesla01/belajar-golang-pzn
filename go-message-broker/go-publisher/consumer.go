package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("consumer Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		fmt.Println("Err Connection :", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Err Channel :", err)
	}
	defer ch.Close()

	msg, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msg {
			fmt.Printf("Receive : %s \n", d.Body)
		}
	}()

	fmt.Println("[*] Waiting Message....")

	<-forever

}
