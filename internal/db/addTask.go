package db

import "gorm.io/gorm"

func AddTask(db *gorm.DB, title string) error {
	task := Task{Title: title}

	return db.Create(&task).Error
}