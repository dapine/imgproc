package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	// conn, err := amqp.Dial("amqp://guest:guest@host.docker.internal:5672")
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

            headers := d.Headers
            payload := d.Body

	    contentType := headers["content_type"].(string)

            switch headers["type"] {
            case "resize":
                width := headers["width"].(int64)
                height := headers["height"].(int64)

		img := Resize(payload, width, height)

                ret := amqp.Publishing{
		    CorrelationId: d.CorrelationId,
                    Timestamp: time.Now(),
                    ContentType: contentType,
                    Body: img,
                }

                err := ch.Publish("", d.ReplyTo, false, false, ret)
                if err != nil {
                    fmt.Println(err)
                }
                break
            default:
                fmt.Println("%+v", d)
            }
		}
	}()

	fmt.Println("Server started")
	<-forever
}
