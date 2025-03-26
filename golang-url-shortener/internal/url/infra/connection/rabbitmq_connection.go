package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConnection() *RabbitMQConnection {
	conn, err := amqp091.Dial("amqp://user:user@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return &RabbitMQConnection{
		Connection: conn,
		Channel:    ch,
	}
}

type RabbitMQConnection struct {
	Connection *amqp091.Connection
	Channel    *amqp091.Channel
}

func (r *RabbitMQConnection) PublisherConfig(nameQueue string) error {
	err := r.Channel.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}

	q, err := r.Channel.QueueDeclare(
		"logs", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		return err
	}

	err = r.Channel.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQConnection) PublishMsg(msg []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Channel.PublishWithContext(ctx,
		"logs",
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	if err != nil {
		fmt.Println("error publish")
		return err
	}

	log.Printf(" [x] Sent %s", msg)
	return nil
}

func (r *RabbitMQConnection) Close() {
	r.Connection.Close()
	r.Channel.Close()
}
