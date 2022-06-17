package rabbitmq

import (
	"encoding/json"
	"fmt"
	delivery "github.com/aliaydins/go-rabbitmq-example/services.delivery/src/internal"
	"github.com/streadway/amqp"
)

func (r *RabbitMQ) NewConsumer(s *delivery.Service) {

	go func() {
		fmt.Println("closing: %s", <-r.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	queue, err := r.channel.QueueDeclare(
		"delivery-queue", // name of the queue
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // noWait
		nil,              // arguments
	)
	if err != nil {
		fmt.Printf("Queue Declare: %s", err)
		return
	}
	fmt.Println("Queue Declare with name delivery-queue-> OK")

	if err = r.channel.QueueBind(
		queue.Name,       // name of the queue
		"",               // bindingKey
		"order-exchange", // sourceExchange
		false,            // noWait
		nil,              // arguments
	); err != nil {
		fmt.Printf("Queue Bind: %s", err)
		return
	}

	fmt.Println("Queue Bind to order-exchange -> OK")

	deliveries, err := r.channel.Consume(
		queue.Name,          // name
		"delivery-consumer", // consumerTag,
		false,               // noAck
		false,               // exclusive
		false,               // noLocal
		false,               // noWait
		nil,                 // arguments
	)
	if err != nil {
		fmt.Printf("Queue Consume: %s", err)
		return
	}

	fmt.Println("Queue Consume with delivery-consumer -> OK")

	for {
		message := <-deliveries

		var eventBody OrderEventType
		err := json.Unmarshal(message.Body, &eventBody)

		fmt.Println(eventBody.ID, eventBody.Name, eventBody.Status, "Consumed from order-exchange")

		if err != nil {
			fmt.Println("Can;t unmarshal the byte array")
			return
		}
		createDelivery := delivery.CreateDeliveryRequest{Name: eventBody.Name}
		d, err := s.CreateDelivery(&createDelivery)

		deliveryCreatedEvent := DeliveryEventType{
			ID:     eventBody.ID,
			Done:   d.Done,
			Status: 1,
		}

		if err != nil {
			deliveryCreatedEvent.Status = 2
			failPayload, _ := json.Marshal(deliveryCreatedEvent)
			r.Publish(failPayload)
			return
		}
		successPayload, _ := json.Marshal(deliveryCreatedEvent)
		r.Publish(successPayload)

		message.Ack(false)

	}
}
