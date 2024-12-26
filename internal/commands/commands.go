package commands

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

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
	switch update.Message.Command() {
	case HelpCmd:
		msg.Text = "У меня есть команды /addTask, /deleteTask и /update."
	case AddTaskCmd:
		handleAddTask(update, msg)
	case DeleteTaskCmd:
		msg.Text = "Какую задачу Вы хотите удалить?"
	case UpdateCmd:
		msg.Text = "Не могу обновить Ваши задачи"
	case TasksList:
		// tasks, err := db.GetAllTasks(dbInstance)

		// if err != nil {
		// 	log.Printf("Error retrieving tasks: %v", err)
		// 	msg.Text = "Ошибка при получении списка задач."
		// 	return
		// }

		taskList := "Ваши задачи:\n"
		// for _, task := range tasks {
		// 	taskList += fmt.Sprintf("- %s\n", task.Title)
		// }
		msg.Text = taskList
	default:
		msg.Text = "Такой команды у меня нет :("
	}
}

func handleAddTask(update tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	taskTitle := update.Message.CommandArguments()

	if taskTitle == "" {
		msg.Text = "Введите название задачи после команды. К примеру: /add Купить молоко"
		return
	}

	task := map[string]string{"title": taskTitle}
	taskJSON, err := json.Marshal(task)
	if err != nil {
		log.Printf("Error marshalling task data: %v", err)
		msg.Text = "Failed to create the task."
		return
	}

	resp, err := http.Post("http://localhost:4000/api/task", "application/json", bytes.NewBuffer(taskJSON))
	if err != nil {
		log.Printf("Error sending POST request: %v", err)
		msg.Text = "Error adding the task."
		return
	}
	defer resp.Body.Close()

	if resp.Status == "200" {
		msg.Text = "Task added successfully!"
	} else {
		msg.Text = "Failed to add the task. Please try again later."
	}
}
