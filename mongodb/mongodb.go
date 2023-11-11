package mongodb

import (
    "os"
    "log"
    "context"
    "time"
    "encoding/json"

    . "github.com/R0L3eX9/softbinator-hackathon/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func DBRead() ([]Category, error) {
	// Set up MongoDB connection options
    DB_URI := os.Getenv("DB_URI")
	clientOptions := options.Client().ApplyURI(DB_URI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
        return nil, err
	}
	defer client.Disconnect(context.Background())

	// Ping the MongoDB server to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
        log.Println(err)
        return nil, err
	}
	log.Println("Connected to MongoDB!")

	// Specify the database and collection
	database := client.Database("api-db")
	collection := database.Collection("categories")

	// Define a filter (optional) to query specific documents

	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find documents in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
        log.Println(err)
        return nil, err
	}
	defer cursor.Close(ctx)

    var results []bson.M
    if err = cursor.All(context.TODO(), &results); err != nil {
        log.Println(err)
        return nil, err
    }
    jsonResult, _ := json.MarshalIndent(&results, "", " ")

    var categories []Category
    err = json.Unmarshal(jsonResult, &categories)
    return categories, nil
}

func AddUserRoadmap(roadmap Roadmap) error {
	// Set up MongoDB connection options
    DB_URI := os.Getenv("DB_URI")
	clientOptions := options.Client().ApplyURI(DB_URI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
        return err
	}
	defer client.Disconnect(context.Background())

	// Ping the MongoDB server to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
        log.Println(err)
        return err
	}
	log.Println("Connected to MongoDB!")

	// Specify the database and collection
	database := client.Database("api-db")
	collection := database.Collection("user_categories")

	// Define a filter (optional) to query specific documents

	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

    _, err = collection.InsertOne(ctx, roadmap)
    if err != nil {
        return err;
    }
    return nil
}
