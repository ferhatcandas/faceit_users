package producer

import (
	"context"
	"faceit_users/pkg/mongo"
	"faceit_users/pkg/rabbitmq"
	"time"

	"github.com/streadway/amqp"
)

func Produce(broker rabbitmq.MessageBroker, eventsRepository mongo.EventsRepository) error {
	for {
		events, err := eventsRepository.GetAll()
		if err != nil {
			return err
		}
		for _, value := range events {
			broker.ExchangeDeclare(value.Exchange, amqp.ExchangeTopic)
			err := broker.Publish(value.Exchange, value.RoutingType, []byte(value.Payload))
			if err == nil {
				err = eventsRepository.DeleteOne(context.Background(), value.Id)
				if err != nil {
					return err
				}
			}
		}
		time.Sleep(time.Millisecond * 300)
	}
}
