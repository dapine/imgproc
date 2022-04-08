package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func processMessages(ch *amqp.Channel, msgs <-chan amqp.Delivery) {
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
		logErr(err)
		break
		default:
		fmt.Println("%+v", d)
		}
	}
}

func main() {
	queue, exchage, key := "hello", "hello_exchange", "teste"

	conn := newConnection("amqp://guest:guest@localhost:5672")
	defer conn.Close()
	ch := newChannel(conn)
	defer ch.Close()

	q := newQueue(ch, queue, exchage, key)

	msgs := consume(ch, q)

	forever := make(chan bool)

	go processMessages(ch, msgs)

	fmt.Println("Server started")
	<-forever
}

