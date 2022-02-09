package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoClient struct {
	client *mongo.Client
	Ping   func(ctx context.Context, rp *readpref.ReadPref) error
}

func (m *mongoClient) UseSession(ctx context.Context, fn func(sc mongo.SessionContext) error) error {
	return m.client.UseSessionWithOptions(ctx, nil, func(sc mongo.SessionContext) error {
		// sc.StartTransaction() -- > for mongo replicaset
		err := fn(sc)
		if err != nil {
			// sc.AbortTransaction(ctx) -- > for mongo replicaset
			return err
		}
		// if err := sc.CommitTransaction(sc); err != nil { -- > for mongo replicaset
		// 	sc.AbortTransaction(ctx)
		// 	panic(err)
		// }
		return nil
	})

}

type MongoRepository interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

func NewMongoClient(connectionString string) *mongoClient {
	clientOptions := options.Client().ApplyURI(connectionString)

	mClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = mClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return &mongoClient{client: mClient, Ping: mClient.Ping}

}
