package rabbitmq

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type MessageBroker struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	consumers  []func()
}

func Connect(connectionStr string) MessageBroker {
	connection, err := amqp.Dial(connectionStr)
	if err != nil {
		panic(err)
	}
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	return MessageBroker{connection: connection, channel: channel}
}

func (m *MessageBroker) ExchangeDeclare(name, exType string) {
	err := m.channel.ExchangeDeclare(
		name,   // name
		exType, // type
		true,   // durable
		false,  // auto-deleted
		false,  // internal
		false,  // noWait
		nil,    // arguments
	)
	if err != nil {
		panic(err)
	}
}

func (m *MessageBroker) QueueDeclare(name, bindExchange, bindRoutingKey string) {
	m.channel.QueueDeclare(
		name,  // name, leave empty to generate a unique name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if bindExchange != "" {
		err := m.channel.QueueBind(
			name,           // name of the queue
			bindRoutingKey, // bindingKey
			bindExchange,   // sourceExchange
			false,          // noWait
			nil,            // arguments
		)

		if err != nil {
			fmt.Println(err)
		}
	}

}

func (m *MessageBroker) RunConsumers() error {
	forever := make(chan bool)

	for _, consumer := range m.consumers {
		consumer()
	}

	<-forever

	return nil
}

func (m *MessageBroker) AddConsumer(qName, consumerName, exchange, exchangeType, routingKey string, fn func(message string) error) {
	ch, err := m.connection.Channel()
	if err != nil {
		fmt.Println(err)
	} else {
		m.ExchangeDeclare(exchange, exchangeType)
		m.QueueDeclare(qName, exchange, routingKey)
		msgs, _ := ch.Consume(
			qName,        // queue
			consumerName, // consumer
			true,         // auto-ack
			false,        // exclusive
			false,        // no-local
			false,        // no-wait
			nil,          // args
		)

		m.consumers = append(m.consumers, func() {
			go func() {
				for d := range msgs {
					err := fn(string(d.Body))
					if err != nil {
						fmt.Printf("Consumer Name: %s\nError: %s", consumerName, err.Error())
					}
				}
			}()
		})
	}

}

func (m *MessageBroker) Publish(exchangeName, routingKey string, payload []byte) error {

	return m.channel.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Transient,
		ContentType:  "application/json",
		Body:         payload,
		Timestamp:    time.Now(),
	})
}
