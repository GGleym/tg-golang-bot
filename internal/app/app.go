package app

import (
	"github/GGleym/telegram-todo-app-golang/internal/config"
	"github/GGleym/telegram-todo-app-golang/internal/repository"
)

func GetBotManager() *repository.TaskRepository {
	collection := config.GetMongoCollection("Cluster0", "todolist")
	repository := repository.NewTaskRepository(collection)

	return repository
}
