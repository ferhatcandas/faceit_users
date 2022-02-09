package entities

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	CorrelationId string             `bson:"correlationId,omitempty"`
	EventType     string             `bson:"eventType,omitempty"`
	Exchange      string             `bson:"exchange,omitempty"`
	RoutingType   string             `bson:"routingType,omitempty"`
	Payload       string             `bson:"payload,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt,omitempty"`
}

func NewEvent(correlationId, eventType, exchange, routingType string, payload interface{}) Event {
	byteArr, _ := json.Marshal(payload)
	return Event{
		Id:            primitive.NewObjectID(),
		CorrelationId: correlationId,
		Exchange:      exchange + ":" + eventType,
		EventType:     eventType,
		RoutingType:   routingType,
		Payload:       string(byteArr),
		CreatedAt:     time.Now(),
	}
}
