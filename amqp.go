package main

import (
	"github.com/streadway/amqp"
)

func newConnection(amqpURI string) *amqp.Connection {
	conn, err := amqp.Dial(amqpURI)
	logErr(err)

	return conn
}

func newChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	logErr(err)

	return ch
}

func newQueue(ch *amqp.Channel, name, exchangeName, key string) amqp.Queue {
	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	logErr(err)

	err = ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	logErr(err)

	err = ch.QueueBind(name, key, exchangeName, false, nil)
	logErr(err)

	return q
}

func consume(ch *amqp.Channel, queue amqp.Queue) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	logErr(err)

	return msgs
}
