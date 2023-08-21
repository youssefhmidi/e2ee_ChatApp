package models

import (
	"context"

	"gorm.io/gorm"
)

// Main struct
type Message struct {
	gorm.Model
	UserID     uint
	ChatRoomID uint
	Message    string
}

// Additionnal struct
type MessageRequest struct {
	Message string `json:"message"`
	Reciver string `json:"reciver"`
}

// Intefaces
type MessageRepository interface {
	// Create a Message
	CreateMessage(ctx context.Context, message Message) error

	// get a Message
	GetMessageByID(ctx context.Context, ID uint) (Message, error)
	// Fetching multiple Messages

	GetMsgsFromUser(ctx context.Context, limit int, user User) ([]Message, error)
	GetMsgsFromRoom(ctx context.Context, limit int, chatRoom ChatRoom) ([]Message, error)

	// Update a Message
	UpdateMessage(ctx context.Context, message Message, target string, value interface{}) error

	// Delete a Message
	DeleteMessage(ctx context.Context, message Message) error
}

type ChatService interface {
	// this method verify every message before encrypting and sending it
	VerifyMessage(ctx context.Context, sender User, signedMessage string) error

	// Create a message object in the databse with the provided room
	// and sender arguments ,the actual message sending part is found in the
	// internal package websocket
	SendMessage(ctx context.Context, message Message) error
}
