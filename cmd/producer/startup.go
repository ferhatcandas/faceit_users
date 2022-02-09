package producer

import (
	"faceit_users/internal/producer"
	"faceit_users/internal/producer/models/config"
	confpkg "faceit_users/pkg/config"
	"faceit_users/pkg/mongo"
	"faceit_users/pkg/rabbitmq"
	"fmt"
)

func Execute(args []string) error {
	fmt.Println("producer started")

	var conf config.Config
	err := confpkg.LoadYMLConfig("configs/producer.yaml", &conf)
	if err != nil {
		panic(err)
	}
	conf.SetMongoHost()
	conf.SetRabbitHost()
	broker := rabbitmq.Connect(conf.RabbitMQUri)
	client := mongo.NewMongoClient(conf.ConnectionString)
	eventsRepository := mongo.NewEventsRepository(*client, conf.DBName)

	return producer.Produce(broker, eventsRepository)
}
