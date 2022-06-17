package main

import (
	"fmt"
	config "github.com/aliaydins/go-rabbitmq-example/services.delivery/src/config"
	"github.com/aliaydins/go-rabbitmq-example/services.delivery/src/entity"
	delivery "github.com/aliaydins/go-rabbitmq-example/services.delivery/src/internal"
	"github.com/aliaydins/go-rabbitmq-example/services.delivery/src/rabbitmq"
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

	db.AutoMigrate(&entity.Delivery{})

	repo := delivery.NewRepository(db)
	service := delivery.NewService(repo)

	r, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		fmt.Println("New RabbitMQ Instance is failed")
		return
	}

	r.NewConsumer(service)

}
