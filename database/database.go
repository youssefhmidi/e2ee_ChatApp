package database

import (
	"context"
	"fmt"
	"log"

	"github.com/youssefhmidi/E2E_encryptedConnection/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// interface for simple interaction with the database
type SqliteDatabase interface {
	InitModels()
	// Creates :
	Insert(ctx context.Context, In interface{}) error
	// Gets :
	GetAll(ctx context.Context, limit int, Model interface{}) (interface{}, error)
	GetAllWhere(ctx context.Context, limit int, Model interface{}, Col string, Condition string) (interface{}, error)
	GetModelById(ctx context.Context, Model interface{}, ID uint) (interface{}, error)
	GetModelWhere(ctx context.Context, Model interface{}, Col string, Condition string) (interface{}, error)
	// Updates :
	// AppendTo(ctx context.Context, Model interface{}, Assosiation string, in interface{}) error
	// UpdateModel(ctx context.Context, Model interface{}, col string, value interface{}) error
	// Delete :
	// DeleteModel(ctx context.Context, Model interface{}) error
}

type Database struct {
	Database *gorm.DB
}

// constructor
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

// initializor
func (db *Database) InitModels() {
	log.Println("initilizing models....")
	err := db.Database.AutoMigrate(&models.User{}, &models.Message{}, &models.ChatRoom{})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("initilizing models successfully")
}

/*
	==================
	== SETTING DATA ==
	==================
*/

func (db *Database) Insert(ctx context.Context, in interface{}) error {
	return db.Database.WithContext(ctx).Create(in).Error
}

/*
	==================
	== GETTING DATA ==
	==================
*/

// make sure the Model arg is a pointer so it can store the response value in it
func (db *Database) GetAll(ctx context.Context, limit int, Model interface{}) (interface{}, error) {
	res := db.Database.WithContext(ctx).Limit(limit).Find(Model)
	return Model, res.Error
}

// make sure the Model arg is a pointer so it can store the response value in it
func (db *Database) GetAllWhere(ctx context.Context, limit int, Model interface{}, Col string, Condition string) (interface{}, error) {
	res := db.Database.WithContext(ctx).Limit(limit).Find(Model, fmt.Sprintf("%v = ?", Col), Condition)
	return Model, res.Error
}

// make sure the Model arg is a pointer so it can store the response value in it
func (db *Database) GetModelById(ctx context.Context, Model interface{}, ID uint) (interface{}, error) {
	res := db.Database.WithContext(ctx).First(Model, ID)
	return Model, res.Error
}

// make sure the Model arg is a pointer so it can store the response value in it
func (db *Database) GetModelWhere(ctx context.Context, Model interface{}, Col string, Condition string) (interface{}, error) {
	res := db.Database.WithContext(ctx).First(Model, fmt.Sprintf("%v = ?", Col), Condition)
	return Model, res.Error
}
