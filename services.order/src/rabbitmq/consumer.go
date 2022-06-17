package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func (r *RabbitMQ) NewConsumer() {

	go func() {
		fmt.Println("closing: %s", <-r.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	queue, err := r.channel.QueueDeclare(
		"order-queue", // name of the queue
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // noWait
		nil,           // arguments
	)
	if err != nil {
		fmt.Printf("Queue Declare: %s", err)
		return
	}
	fmt.Println("Queue Declare with name order-queue-> OK")

	if err = r.channel.QueueBind(
		queue.Name,          // name of the queue
		"",                  // bindingKey
		"delivery-exchange", // sourceExchange
		false,               // noWait
		nil,                 // arguments
	); err != nil {
		fmt.Printf("Queue Bind: %s", err)
		return
	}

	fmt.Println("Queue Bind to order-exchange -> OK")

	deliveries, err := r.channel.Consume(
		queue.Name,       // name
		"order-consumer", // consumerTag,
		false,            // noAck
		false,            // exclusive
		false,            // noLocal
		false,            // noWait
		nil,              // arguments
	)
	if err != nil {
		fmt.Printf("Queue Consume: %s", err)
		return
	}

	fmt.Println("Queue Consume with order-consumer -> OK")

	for {
		message := <-deliveries

		var eventBody DeliveryEventType
		err := json.Unmarshal(message.Body, &eventBody)

		fmt.Println(eventBody.ID, eventBody.Done, eventBody.Status, "Consumed from delivery-exchange")

		if err != nil {
			fmt.Println("Can;t unmarshal the byte array")
			return
		}

		/*createOrder := order.UpdateOrderRequest{
			ID:     eventBody.ID,
			Status: eventBody.Status,
		}

		_, err = s.UpdateOrder(&createOrder)
		if err != nil {
			fmt.Println("Order Status update failed with id ->", createOrder.ID)
			return
		}
		fmt.Println("Order Status updated successfully with id ->", createOrder.ID)*/

		message.Ack(false)

	}
}
