package database

import (
	"log"

	"github.com/youssefhmidi/E2E_encryptedConnection/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDatabase interface {
	InitModels()
}

type Database struct {
	Database *gorm.DB
}

func NewDB(location string) SqliteDatabase {
	DB, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	log.Println("Connecting to the database....")
	if err != nil {
		log.Fatal(err)
		return &Database{}
	}
	log.Println("connection has been done successfully")
	return &Database{
		Database: DB,
	}
}

func (db *Database) InitModels() {
	log.Println("initilizing models....")
	err := db.Database.AutoMigrate(&models.User{}, &models.Message{}, &models.ChatRoom{})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("initilizing models successfully")
}
