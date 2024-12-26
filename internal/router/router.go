package router

import (
	"github/GGleym/telegram-todo-app-golang/internal/config"
	"github/GGleym/telegram-todo-app-golang/internal/controller"
	"github/GGleym/telegram-todo-app-golang/internal/repository"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	collection := config.GetMongoCollection("Cluster0", "todolist")
	repo := repository.NewTaskRepository(collection)
	controller := controller.NewTaskController(repo)

	router.HandleFunc("/api/tasks", controller.GetAllTasks).Methods("GET")
	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", controller.MarkAsDone).Methods("PUT")
	router.HandleFunc("/api/task/{id}", controller.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteAllTasks).Methods("DELETE")

	return router
}
