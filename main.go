package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func processResizeMessages(ch *amqp.Channel, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		headers := d.Headers
		payload := d.Body

		contentType := headers["content_type"].(string)

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
	}
}

func consumeResizeQueue(ch *amqp.Channel) {
	queue, exchage, key := "resize", "image_processing", "resize"
	q := newQueue(ch, queue, exchage, key)
	msgs := consume(ch, q)
	go processResizeMessages(ch, msgs)
}

func main() {
	conn := newConnection("amqp://guest:guest@localhost:5672")
	defer conn.Close()
	ch := newChannel(conn)
	defer ch.Close()

	consumeResizeQueue(ch)

	forever := make(chan bool)
	fmt.Println("Server started")
	<-forever
}

