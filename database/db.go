package database

import (
    "backend-assignment/models"
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

    if err != nil {
        panic("Failed to connect database")
    }

    database.AutoMigrate(&models.User{})

    DB = database
}