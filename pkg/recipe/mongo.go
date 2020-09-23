package recipe

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://recipe_db"))
	if err != nil {
		return nil, err
	}
	// From docs: defer a call to Disconnect after instantiating your client
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("could not disconnect: %v", err)
		}
	}()

	// check connection
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	log.Println("Connected to db!!")
	return client, nil
}
