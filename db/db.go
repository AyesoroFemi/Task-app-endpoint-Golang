package db


import (
    "context"
    "os"
    "log"
	"time"
	"github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectToMongo() (*mongo.Client, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading  the .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDb connection string
	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to mongo...")

	return client, nil
}