package main

import (
	"fmt"
	"time"

	"github.com/dapine/imgproc/image"
	"github.com/streadway/amqp"
)

func processResizeMessages(ch *amqp.Channel, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		headers := d.Headers
		payload := d.Body

		contentType := headers["content_type"].(string)

		width := headers["width"].(int64)
		height := headers["height"].(int64)

		img := image.Resize(payload, width, height)

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

func processRotateMessages(ch *amqp.Channel, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		headers := d.Headers
		payload := d.Body

		contentType := headers["content_type"].(string)

		angle := headers["angle"].(int64)

		img := image.Rotate(payload, angle)

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

func consumeQueues(ch *amqp.Channel) {
	queue, exchage, key := "rotate", "image_processing", "rotate"
	qRotate := newQueue(ch, queue, exchage, key)

	queue, exchage, key = "resize", "image_processing", "resize"
	qResize := newQueue(ch, queue, exchage, key)

	msgsRotate := consume(ch, qRotate)
	msgsResize := consume(ch, qResize)

	go processResizeMessages(ch, msgsResize)
	go processRotateMessages(ch, msgsRotate)
}

func main() {
	conn := newConnection("amqp://guest:guest@localhost:5672")
	defer conn.Close()
	ch := newChannel(conn)
	defer ch.Close()

	consumeQueues(ch)

	forever := make(chan bool)
	fmt.Println("Server started")
	<-forever
}

