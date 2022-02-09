package consumer

import (
	"faceit_users/internal/consumer"
	"faceit_users/internal/consumer/models/config"
	confpkg "faceit_users/pkg/config"
	"faceit_users/pkg/rabbitmq"
	"fmt"

	"github.com/streadway/amqp"
)

func Execute(args []string) error {
	fmt.Println("consumer started")

	var conf config.Config
	err := confpkg.LoadYMLConfig("configs/consumer.yaml", &conf)
	if err != nil {
		panic(err)
	}
	conf.SetRabbitHost()
	broker := rabbitmq.Connect(conf.RabbitMQUri)

	broker.AddConsumer("FaceIt.User.Created", "UserCreated", "User:Created", amqp.ExchangeTopic, "", consumer.UserCreatedConsumer)
	broker.AddConsumer("FaceIt.User.Updated", "UserUpdated", "User:Updated", amqp.ExchangeTopic, "", consumer.UserUpdatedConsumer)
	broker.AddConsumer("FaceIt.User.Deleted", "UserDeleted", "User:Deleted", amqp.ExchangeTopic, "", consumer.UserDeletedConsumer)

	return broker.RunConsumers()
}
