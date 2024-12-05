package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
