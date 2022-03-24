package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	conn, err := amqp.Dial("amqp://guest:guest@host.docker.internal:5672")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	err = ch.ExchangeDeclare("hello_exchange", "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = ch.QueueBind("hello", "teste", "hello_exchange", false, nil)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// fmt.Println("New message: %s", string(d.Body))
			content := d.Body

			img, err := bimg.NewImage(content).Resize(32, 32)
			if err != nil {
				fmt.Println(err)
			

			// fmt.Println(d)
			// fmt.Println("New message")
		}
	}()

	fmt.Println("Server started")
	<-forever
}
