package repository

import (
	"context"
	"errors"
	"github/GGleym/telegram-todo-app-golang/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) *TaskRepository {
	return &TaskRepository{collection: collection}
}

func (r *TaskRepository) InsertTask(ctx context.Context, task model.Task) (primitive.ObjectID, error) {
	res, err := r.collection.InsertOne(ctx, task)

	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return primitive.NilObjectID, errors.New("failed to convert inserted ID to ObjectID")
	}

	return id, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"done": true}}
	res, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}
	res, err := r.collection.DeleteOne(ctx, filter)

	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func (r *TaskRepository) DeleteAllTasks(ctx context.Context) (int64, error) {
	res, err := r.collection.DeleteMany(ctx, bson.D{})

	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]bson.M, error) {
	cur, err := r.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var tasks []bson.M

	for cur.Next(ctx) {
		var task bson.M

		if err := cur.Decode(&task); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
