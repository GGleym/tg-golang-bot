package controller

import (
	"context"
	"encoding/json"
	"github/GGleym/telegram-todo-app-golang/internal/model"
	"github/GGleym/telegram-todo-app-golang/internal/repository"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	repo *repository.TaskRepository
}

func NewTaskController(repo *repository.TaskRepository) *TaskController {
	return &TaskController{repo: repo}
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := c.repo.GetAllTasks(context.Background())

	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task model.Task
	
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Incorrect input", http.StatusBadRequest)
		return
	}

	id, err := c.repo.InsertTask(context.Background(), task)

	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	task.ID = id
	json.NewEncoder(w).Encode(task)
}

func (c *TaskController) MarkAsDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		http.Error(w, "Incorrect task ID", http.StatusBadRequest)
		return
	}

	count, err := c.repo.UpdateTask(context.Background(), id)

	if err != nil || count == 0 {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(params)
}

func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		http.Error(w, "Invalid ID for task to delete", http.StatusBadRequest)
		return
	}

	count, err := c.repo.DeleteTask(context.Background(), id)

	if err != nil || count == 0 {
		http.Error(w, "Error while deleting the task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(params["id"])
}

func (c *TaskController) DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	count, err := c.repo.DeleteAllTasks(context.Background())

	if err != nil {
		http.Error(w, "Failed to delete tasks", http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(count)
}
