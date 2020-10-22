package recipe

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	*mongo.Client
}

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"title"`
	Ingredients []string `json:"ingredients"`
}

func NewMongoClient() (Repo, error) {
	// Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Cancel context to avoid memory leak
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@recipe_db")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return Repo{}, err
	}

	// From docs: defer a call to Disconnect after instantiating your client
	// Need to do this from main?
	// defer client.Disconnect(ctx)

	// check connection
	if err = client.Ping(ctx, nil); err != nil {
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
	id := res.InsertedID
	log.Printf("insertedID: %v", id)
	return err
}

func (r Repo) GetRecipesByIngredient(ingredient string) ([]Recipe, error) {
	log.Printf("ingredient %s: ", ingredient)

	collection := r.Client.Database("dev").Collection("recipes")

	findOptions := options.Find()
	// Need to set a limit?
	findOptions.SetLimit(2)

	var allRecipes []Recipe

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return allRecipes, err
	}

	// for multiple docs, a cursor is returned
	// documents are decoded one at a time
	for cursor.Next(context.TODO()) {
		var recipe Recipe

		if err := cursor.Decode(&recipe); err != nil {
			return allRecipes, err
		}

		allRecipes = append(allRecipes, recipe)
	}

	if err := cursor.Err(); err != nil {
		return allRecipes, err
	}
	cursor.Close(context.TODO())

	log.Printf("Found documents: %+v\n", allRecipes)
	return allRecipes, nil
}
