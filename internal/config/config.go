package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetMongoCollection(dbName, colName string) *mongo.Collection {
	LoadEnv()

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

	return client.Database(dbName).Collection(colName)
}
