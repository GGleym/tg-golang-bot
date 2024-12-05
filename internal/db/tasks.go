package db

import "gorm.io/gorm"

func GetAllTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks).Error
	return tasks, err
}