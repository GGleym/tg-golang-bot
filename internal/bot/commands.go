package bot

import (
	"context"
	"fmt"
	"github/GGleym/telegram-todo-app-golang/internal/app"
	"github/GGleym/telegram-todo-app-golang/internal/model"
	"github/GGleym/telegram-todo-app-golang/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	HelpCmd       = "help"
	AddTaskCmd    = "add"
	DeleteTaskCmd = "delete"
	UpdateCmd     = "update"
	TasksList     = "tasks"
)

func HandleCommands(update tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	botManager := app.GetBotManager()

	switch update.Message.Command() {
	case HelpCmd:
		msg.Text = "У меня есть команды /addTask, /deleteTask и /update."
	case AddTaskCmd:
		handleAddTask(update, msg, botManager)
	case DeleteTaskCmd:
		msg.Text = "Какую задачу Вы хотите удалить?"
	case UpdateCmd:
		msg.Text = "Не могу обновить Ваши задачи"
	case TasksList:
		handleTasksShow(msg, botManager)
	default:
		msg.Text = "Такой команды у меня нет :("
	}
}

func handleTasksShow(msg *tgbotapi.MessageConfig, botManager *repository.TaskRepository) {
	tasks, err := botManager.GetAllTasks(context.TODO())

	taskList := "Ваши задачи: \n"

	if err != nil {
		msg.Text = "Не смогли получить Ваши задачи. Попробуйте еще раз."
		return
	}

	for _, task := range tasks {
		done := task["done"]
		taskDescription := task["task"]

		taskList += fmt.Sprintf("- Сделанная: %v, Задача: %v\n", done, taskDescription)
	}

	msg.Text = taskList
}

func handleAddTask(update tgbotapi.Update, msg *tgbotapi.MessageConfig, botManager *repository.TaskRepository) {
	taskTitle := update.Message.CommandArguments()

	if taskTitle == "" {
		msg.Text = "Введите название задачи после команды. К примеру: /add Купить молоко"
		return
	}

	// task := map[string]string{"title": taskTitle}
	// taskJSON, err := json.Marshal(task)

	// if err != nil {
	// 	log.Printf("Error marshalling task data: %v", err)
	// 	msg.Text = "Failed to create the task."
	// 	return
	// }

	if _, err := botManager.InsertTask(context.TODO(), model.Task{Task: taskTitle}); err != nil {
		msg.Text = "Не смогли добавить Вашу задачу. Попробуй еще раз."
		return
	}

	msg.Text = fmt.Sprintf("Добавили вашу задачу: %s", taskTitle)
}
