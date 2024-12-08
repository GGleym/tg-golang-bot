package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github/GGleym/telegram-todo-app-golang/internal/db/model"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "Cluster0"
const colName = "todolist"

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

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	collection = client.Database(dbName).Collection(colName)

	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	// 	panic(err)
	// }

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func insertTask(task model.Task) {
	fmt.Printf("Data inserted: %+v", task)
	inserted, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal("Error while inserting the task: ", err)
	}

	fmt.Println("Inserted the: ", inserted.InsertedID)
}

func updateTask(taskId string) {
	id, _ := primitive.ObjectIDFromHex(taskId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"done": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal("Error updating the task: ", err)
	}

	fmt.Println("modified count: ", res.ModifiedCount)
}

func deleteTask(taskId string) {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal("Error deleting task: ", err)
	}

	fmt.Println("deleted task: ", res.DeletedCount)
}

func deleteAllTasks() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal("Error deleting items", err)
	}

	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

func getAllTasks() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal("Error getting all tasks: ", err)
	}

	var tasks []primitive.M

	for cur.Next(context.Background()) {
		var task bson.M

		err := cur.Decode(&task)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	defer cur.Close(context.Background())

	return tasks
}

func GetAllTasksReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-ulrencode")

	allTasks := getAllTasks()
	json.NewEncoder(w).Encode(allTasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var task model.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	insertTask(task)
	json.NewEncoder(w).Encode(task)
}

func MaskAsDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateTask(params["id"])
	json.NewEncoder(w).Encode(params)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllTasks()
	json.NewEncoder(w).Encode(count)
}
