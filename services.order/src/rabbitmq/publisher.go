package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func (r RabbitMQ) Publish(body []byte) {

	if err := r.channel.Publish(
		"order-exchange", // publish to an exchange
		"",               // routing to 0 or more queues
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		fmt.Printf("Exchange Publish: %s", err)
	}

}
