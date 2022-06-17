package main

import (
	"fmt"

	"github.com/aliaydins/go-rabbitmq-example/services.order/src/config"
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/entity"
	order "github.com/aliaydins/go-rabbitmq-example/services.order/src/internal"
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/rabbitmq"
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := config.LoadConfig(".")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the DB.")
	}

	db.AutoMigrate(&entity.Order{})

	r, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		fmt.Println("New RabbitMQ Instance is failed")
		return
	}
	repo := order.NewRepository(db, r)
	service := order.NewService(repo, r)
	handler := order.NewHandler(service)

	go r.NewConsumer()

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}

}
