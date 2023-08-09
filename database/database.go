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
	// init all models
	InitModels()
	// Creates :

	// make sure the in arg is a pointer
	Insert(ctx context.Context, In interface{}) error
	// Gets :

	// make sure the Model arg is a pointer so it can store the response value in it
	GetAll(ctx context.Context, limit int, Model interface{}) (interface{}, error)
	// make sure the Model arg is a pointer so it can store the response value in it
	GetAllWhere(ctx context.Context, limit int, Model interface{}, Col string, Condition interface{}) (interface{}, error)
	// make sure the Model arg is a pointer so it can store the response value in it
	GetModelById(ctx context.Context, Model interface{}, ID uint) (interface{}, error)
	// make sure the Model arg is a pointer so it can store the response value in it
	GetModelWhere(ctx context.Context, Model interface{}, Col string, Condition interface{}) (interface{}, error)

	// Updates :

	// appending to an assosiation a field in the model struct type that refer to another struct,
	// E.g: the Messages field in User struct is an association
	//
	// ❗NOTE: Make sure the Model arg is a pointer❗ (I know goofy emoji)
	AppendTo(ctx context.Context, Model interface{}, Assosiation string, in interface{}) error
	// Make sure the Model is not empty otherwise use UpdateWhere to Update a Model with a condition
	UpdateModel(ctx context.Context, Model interface{}, col string, value interface{}) error
	//	 // The ModelType argument take an empty struct of the type
	//		// E.g : if you want to update a user's name and you have only one Information about the user
	//		UpdateWhere(c, &User{}, "active", true, "name", "active user")
	//		// same as : UPDATE user SET name="active user" WHERE active=true
	//
	//		this a bad exmaple but I guess you get the point of this function now.
	UpdateWhere(ctx context.Context, ModelType interface{}, condition_col string, condition_val interface{}, col string, value interface{}) error

	// Delete :

	// Make sure you're not passing an empty model To get a model use GetModelById() or GetModelWhere() methods
	DeleteModel(ctx context.Context, Model interface{}) error
}

type Database struct {
	// A database struct mainly for simplictiy
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

// make sure the in arg is a pointer
func (db *Database) Insert(ctx context.Context, in interface{}) error {
	return db.Database.WithContext(ctx).Create(in).Error
}

/*
	==================
	== GETTING DATA ==
	==================
*/

// make sure the Model arg is a pointer and a slice type so it can store the response value in it
func (db *Database) GetAll(ctx context.Context, limit int, Model interface{}) (interface{}, error) {
	res := db.Database.WithContext(ctx).Limit(limit).Find(Model)
	return Model, res.Error
}

// make sure the Model arg is a pointer so it can store the response value in it
func (db *Database) GetAllWhere(ctx context.Context, limit int, Model interface{}, Col string, Condition interface{}) (interface{}, error) {
	res := db.Database.WithContext(ctx).Limit(limit).Find(Model, fmt.Sprintf("%v = ?", Col), Condition)
	return Model, res.Error
}

// make sure the Model arg is a pointer so it can store the response value in it
func (db *Database) GetModelById(ctx context.Context, Model interface{}, ID uint) (interface{}, error) {
	res := db.Database.WithContext(ctx).First(Model, ID)
	return Model, res.Error
}

// make sure the Model arg is a pointer so it can store the response value in it
func (db *Database) GetModelWhere(ctx context.Context, Model interface{}, Col string, Condition interface{}) (interface{}, error) {
	res := db.Database.WithContext(ctx).First(Model, fmt.Sprintf("%v = ?", Col), Condition)
	return Model, res.Error
}

/*
	===================
	== UPDATING DATA ==
	===================
*/

// appending to an assosiation a field in the model struct type that refer to another struct,
// E.g: the Messages field in User struct is an association
//
// ❗NOTE: Make sure the Model arg is a pointer❗ (I know goofy emoji)
func (db *Database) AppendTo(ctx context.Context, Model interface{}, Assosiation string, in interface{}) error {
	return db.Database.Model(Model).WithContext(ctx).Association(Assosiation).Append(in)
}

// Make sure the Model is not empty otherwise use UpdateWhere to Update a Model with a condition
func (db *Database) UpdateModel(ctx context.Context, Model interface{}, col string, value interface{}) error {
	return db.Database.Model(Model).WithContext(ctx).Update(col, value).Error
}

//	 // The ModelType argument take an empty struct of the type
//		// E.g : if you want to update a user's name and you have only one Information about the user
//		UpdateWhere(c, &User{}, "active", true, "name", "active user")
//		// same as : UPDATE user SET name="active user" WHERE active=true
//
//		this a bad exmaple but I guess you get the point of this function now.
func (db *Database) UpdateWhere(ctx context.Context, ModelType interface{}, condition_col string, condition_val interface{}, col string, value interface{}) error {
	return db.Database.Model(ModelType).WithContext(ctx).Where(fmt.Sprintf("%v = ?", condition_col), condition_val).Update(col, value).Error
}

/*
	===================
	== DELETING DATA ==
	===================
*/

// Make sure you're not passing an empty model To get a model use GetModelById() or GetModelWhere() methods
func (db *Database) DeleteModel(ctx context.Context, Model interface{}) error {
	return db.Database.Delete(Model).Error
}
