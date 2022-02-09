package mongo

import (
	"context"
	"faceit_users/pkg/mongo/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	mongoClient *mongoClient
	repository  MongoRepository
}

func NewUserRepository(client mongoClient, database string) UserRepository {
	return UserRepository{
		repository:  client.client.Database(database).Collection("users"),
		mongoClient: &client,
	}
}
func (u *UserRepository) UseSession(ctx context.Context, fn func(sc mongo.SessionContext) error) error {
	return u.mongoClient.UseSession(ctx, fn)
}
func (u *UserRepository) InsertOne(ctx context.Context, user entities.User) error {
	_, err := u.repository.InsertOne(ctx, user, &options.InsertOneOptions{})
	return err
}

func (u *UserRepository) UpdateOne(ctx context.Context, id string, user entities.User) error {
	_, err := u.repository.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

func (u *UserRepository) DeleteOne(ctx context.Context, id string) error {
	_, err := u.repository.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (u *UserRepository) Get(ctx context.Context, country string, pageIndex, pageSize int64) ([]entities.User, error) {
	cursor, err := u.repository.Find(ctx, bson.M{"country": country}, &options.FindOptions{
		Skip:  &pageIndex,
		Limit: &pageSize,
	})
	if err != nil {
		return nil, err
	}
	var users []entities.User
	err = cursor.All(ctx, &users)
	defer cursor.Close(ctx)
	return users, err
}
func (u *UserRepository) FindOne(ctx context.Context, id string) (*entities.User, error) {
	cursor, err := u.repository.Find(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	if !cursor.Next(ctx) {
		return nil, nil
	}
	var user entities.User
	err = cursor.Decode(&user)
	defer cursor.Close(ctx)
	return &user, err
}

func (u *UserRepository) FindOneByNickname(ctx context.Context, nick string) (*entities.User, error) {
	cursor, err := u.repository.Find(ctx, bson.M{"nickname": nick})
	if err != nil {
		return nil, err
	}
	if !cursor.Next(ctx) {
		return nil, nil
	}
	var user entities.User
	err = cursor.Decode(&user)
	defer cursor.Close(ctx)
	return &user, err
}
