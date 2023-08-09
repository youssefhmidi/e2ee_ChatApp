package models

import (
	"context"

	"gorm.io/gorm"
)

type ChatRoom struct {
	gorm.Model
	OwnerID   uint
	Name      string
	IsPublic  bool
	PublicKey string
	Members   []User    `gorm:"many2many:user_chatroom"`
	Messages  []Message `gorm:"foreignKey:ChatRoomID"`
}

type ChatRoomRepository interface {
	// Creat a ChatRoom
	CreateChatRoom(ctx context.Context, room ChatRoom) error

	// Get a ChatRoom
	GetRoomByID(ctx context.Context, ID uint) (ChatRoom, error)
	GetRoomByName(ctx context.Context, Name string) (ChatRoom, error)

	// fetching multiple ChatRoom
	GetRoomsFromUser(ctx context.Context, user User) ([]ChatRoom, error)
	GetOwnedRooms(ctx context.Context, user User) ([]ChatRoom, error)

	// Updating a ChatRoom
	UpdateRoom(ctx context.Context, room ChatRoom) error

	// Delete a ChatRoom
	DeleteRoom(ctx context.Context, room ChatRoom) error
}
