package rabbitmq

import (
	"fmt"
	config "github.com/aliaydins/go-rabbitmq-example/services.delivery/src/config"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
	config.LoadConfig(".")
	user := config.AppConfig.RabbitUser
	pass := config.AppConfig.RabbitPassword
	host := config.AppConfig.RabbitHost
	port := config.AppConfig.RabbitPort

	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port + "/")
	if err != nil {
		fmt.Println(fmt.Errorf("dial: %s", err))
		return nil, err
	}
	fmt.Println("AMQP connection -> OK")

	channel, err := conn.Channel()
	if err != nil {
		fmt.Printf("Channel error:%s", err)
		return nil, err
	}
	fmt.Println("Channel connection -> OK")

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
	}, nil
}
