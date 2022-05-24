package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dapine/imgproc/image"
	"github.com/streadway/amqp"
)

type AmqpImageTransformation func([]byte, amqp.Table) []byte

func AmqpResize(bytes []byte, headers amqp.Table) []byte {
	width := headers["width"].(int64)
	height := headers["height"].(int64)
	return image.Resize(bytes, width, height)
}

func AmqpRotate(bytes []byte, headers amqp.Table) []byte {
	angle := headers["angle"].(int64)
	return image.Rotate(bytes, angle)
}

func AmqpConvert(bytes []byte, headers amqp.Table) []byte {
	it := headers["target_image_type"].(string)
	img, _, _ := image.Convert(bytes, it)
	return img
}

func AmqpCrop(bytes []byte, headers amqp.Table) []byte {
	width := headers["width"].(int64)
	height := headers["height"].(int64)
	gravity := headers["gravity"].(string)
	return image.Crop(bytes, width, height, gravity)
}

func AmqpEnlarge(bytes []byte, headers amqp.Table) []byte {
	width := headers["width"].(int64)
	height := headers["height"].(int64)
	return image.Enlarge(bytes, width, height)
}

func AmqpExtract(bytes []byte, headers amqp.Table) []byte {
	x := headers["x"].(int64)
	y := headers["y"].(int64)
	width := headers["width"].(int64)
	height := headers["height"].(int64)
	return image.Extract(bytes, x, y, width, height)
}

func AmqpFlip(bytes []byte, headers amqp.Table) []byte {
	axis := headers["axis"].(string)
	return image.Flip(bytes, axis)
}

func processMessages(ch *amqp.Channel, msgs <-chan amqp.Delivery, transform AmqpImageTransformation) {
	for d := range msgs {
		headers := d.Headers
		payload := d.Body

		contentType := headers["content_type"].(string)

		img := transform(payload, headers)

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

	queue, exchage, key = "convert", "image_processing", "convert"
	qConvert := newQueue(ch, queue, exchage, key)

	queue, exchage, key = "crop", "image_processing", "crop"
	qCrop := newQueue(ch, queue, exchage, key)

	queue, exchage, key = "enlarge", "image_processing", "enlarge"
	qEnlarge := newQueue(ch, queue, exchage, key)

	queue, exchage, key = "extract", "image_processing", "extract"
	qExtract := newQueue(ch, queue, exchage, key)

	queue, exchage, key = "flip", "image_processing", "flip"
	qFlip := newQueue(ch, queue, exchage, key)

	msgsRotate := consume(ch, qRotate)
	msgsResize := consume(ch, qResize)
	msgsConvert := consume(ch, qConvert)
	msgsCrop := consume(ch, qCrop)
	msgsEnlarge := consume(ch, qEnlarge)
	msgsExtract := consume(ch, qExtract)
	msgsFlip := consume(ch, qFlip)

	go processMessages(ch, msgsRotate, AmqpRotate)
	go processMessages(ch, msgsResize, AmqpResize)
	go processMessages(ch, msgsConvert, AmqpConvert)
	go processMessages(ch, msgsCrop, AmqpCrop)
	go processMessages(ch, msgsEnlarge, AmqpEnlarge)
	go processMessages(ch, msgsExtract, AmqpExtract)
	go processMessages(ch, msgsFlip, AmqpFlip)
}

func main() {
	amqpHostname := os.Getenv("AMQP_HOSTNAME")

	if amqpHostname == "" {
		amqpHostname = "localhost"
	}

	conn := newConnection("amqp://guest:guest@"+amqpHostname+":5672")
	defer conn.Close()
	ch := newChannel(conn)
	defer ch.Close()

	consumeQueues(ch)

	forever := make(chan bool)
	fmt.Println("Server started")
	<-forever
}

