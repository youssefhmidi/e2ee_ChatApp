package models

import (
	"context"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID     uint
	ChatRoomID uint
	Message    string
}

type MessageRepository interface {
	// Create a Message
	CreateMessage(ctx context.Context, message Message, chatroom ChatRoom, user User) error

	// get a Message
	GetMessageByID(ctx context.Context, ID uint) (Message, error)
	// Fetching multiple Messages

	GetMsgsFromUser(ctx context.Context, limit int, user User) ([]Message, error)
	GetMsgsFromRoom(ctx context.Context, limit int, chatRoom ChatRoom) ([]Message, error)

	// Update a Message
	UpdateMessage(ctx context.Context, message Message) error

	// Delete a Message
	DeleteMessage(ctx context.Context, message Message) error
}
