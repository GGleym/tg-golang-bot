package router

import (
	"github/GGleym/telegram-todo-app-golang/internal/db/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", controller.GetAllTasksReq).Methods("GET")
	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", controller.MaskAsDone).Methods("PUT")
	router.HandleFunc("/api/task/{id}", controller.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteAllTasks).Methods("DELETE")

	return router
}
