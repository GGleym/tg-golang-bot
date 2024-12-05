package commands

import (
	"fmt"
	"github/GGleym/telegram-todo-app-golang/internal/db"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

const (
	HelpCmd 	  = "help"
	AddTaskCmd    = "add"
	DeleteTaskCmd = "delete"
	UpdateCmd 	  = "update"
	TasksList     = "tasks"
)

func HandleCommands(update tgbotapi.Update, msg *tgbotapi.MessageConfig, dbInstance *gorm.DB) {
	switch update.Message.Command() {
	case HelpCmd:
		msg.Text = "У меня есть команды /addTask, /deleteTask и /update."
	case AddTaskCmd:
		handleAddTask(update, msg, dbInstance)
	case DeleteTaskCmd:
		msg.Text = "Какую задачу Вы хотите удалить?"
	case UpdateCmd:
		msg.Text = "Не могу обновить Ваши задачи"
	case TasksList:
		tasks, err := db.GetAllTasks(dbInstance)

		if err != nil {
			log.Printf("Error retrieving tasks: %v", err)
			msg.Text = "Ошибка при получении списка задач."
			return
		}

		taskList := "Ваши задачи:\n"
		for _, task := range tasks {
			taskList += fmt.Sprintf("- %s\n", task.Title)
		}
		msg.Text = taskList
	default:
		msg.Text = "Такой команды у меня нет :("
	}
}

func handleAddTask(update tgbotapi.Update, msg *tgbotapi.MessageConfig, dbInstance *gorm.DB) {
	taskTitle := update.Message.CommandArguments()

	if taskTitle == "" {
		msg.Text = "Введите название задачи после команды."
		return
	}

	err := db.AddTask(dbInstance, taskTitle)
	if err != nil {
		log.Printf("Error adding task %v", err)
		msg.Text = "Не удалось добавить задание."
	} else {
		msg.Text = "Задача успешно добавлена!"
	}
}
