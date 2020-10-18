package recipe

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repo struct {
	*mongo.Client
}

func NewMongoClient() (Repo, error) {
	// Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Cancel context to avoid memory leak
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@recipe_db"))
	if err != nil {
		return Repo{}, err
	}

	// From docs: defer a call to Disconnect after instantiating your client
	// Need to do this from main?
	// defer client.Disconnect(ctx)

	// check connection
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return Repo{}, err
	}
	log.Println("Connected to db!!")
	return Repo{client}, nil
}

func (r Repo) AddNumbers() error {
	collection := r.Client.Database("testing").Collection("numbers")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	res, err := collection.InsertOne(context.TODO(), bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		return err
	}

	id := res.InsertedID
	log.Printf("insertedID: %v", id)
	return nil
}
