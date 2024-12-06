package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "admin"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	if err := godotenv.Load(); err != nil {
	  log.Fatal("Error loading .env file")
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword == "" {
		log.Fatal("Environment variable DB_PASSWORD is not set")
	}

	connectionString := fmt.Sprintf(
		"mongodb+srv://gleymscoot:%s@cluster0.ij2hy.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0",
		dbPassword,
	)

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success!")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection = client.Database(dbName).Collection(colName)

	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	// 	panic(err)
	// }

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}