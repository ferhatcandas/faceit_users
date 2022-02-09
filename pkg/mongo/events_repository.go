package mongo

import (
	"context"
	"faceit_users/pkg/mongo/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventsRepository struct {
	mongoClient *mongoClient
	repository  MongoRepository
}

func NewEventsRepository(client mongoClient, database string) EventsRepository {
	return EventsRepository{
		repository:  client.client.Database(database).Collection("events"),
		mongoClient: &client,
	}
}
func NewEventsHistoryRepository(client mongoClient, database string) EventsRepository {
	return EventsRepository{
		repository:  client.client.Database(database).Collection("eventhistories"),
		mongoClient: &client,
	}
}
func (u *EventsRepository) InsertOne(ctx context.Context, user entities.Event) error {
	_, err := u.repository.InsertOne(ctx, user, &options.InsertOneOptions{})
	return err
}

func (u *EventsRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	_, err := u.repository.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
func (u *EventsRepository) GetAll() ([]entities.Event, error) {
	ctx := context.Background()
	cursor, err := u.repository.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var events []entities.Event
	err = cursor.All(ctx, &events)
	defer cursor.Close(ctx)
	return events, err
}
