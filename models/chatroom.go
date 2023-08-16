package models

import (
	"context"

	"gorm.io/gorm"
)

// types
type ChatType string

type ChatRoom struct {
	gorm.Model
	OwnerID   uint
	Name      string
	IsPublic  bool
	PublicKey string
	// Type can be either a "dm" or "group"
	Type     ChatType
	Members  []User    `gorm:"many2many:user_chatroom"`
	Messages []Message `gorm:"foreignKey:ChatRoomID"`
}

type ChatRoomRepository interface {
	// Creat a ChatRoom
	CreateChatRoom(ctx context.Context, room ChatRoom) error

	// Get a ChatRoom
	GetRoomByID(ctx context.Context, ID uint) (ChatRoom, error)
	GetRoomByName(ctx context.Context, Name string) (ChatRoom, error)

	// fetching multiple ChatRoom
	GetRoomsFromUser(ctx context.Context, limit int, user User) ([]ChatRoom, error)
	GetOwnedRooms(ctx context.Context, limit int, user User) ([]ChatRoom, error)

	// Updating a ChatRoom
	UpdateRoom(ctx context.Context, room ChatRoom, target string, value interface{}) error
	AppendToRoom(ctx context.Context, room ChatRoom, association string, in interface{}) error

	// Delete a ChatRoom
	DeleteRoom(ctx context.Context, room ChatRoom) error
}

type ChatRoomService interface {
	// Creates Group and initilize a symetric encryption key
	// returns a key in a pem format and a error (which is nil if the operation doe correctly)
	CreateGroup(ctx context.Context, Owner User, Members []User, IsPublic bool) (string, error)
	// Creates a DM with another user.
	// user1 and user2 are the participent in the DM
	// returns an error if not succeded
	CreateDM(ctx context.Context, user1 User, user2 User) error

	// Add a member to the provided room
	// returns an error if failed
	AddMember(ctx context.Context, Room ChatRoom, user User) error
	// Removes a member from the provided room
	// returns an error if failed
	RemoveMember(ctx context.Context, Room ChatRoom, user User) error

	// Gets all ChatRooms with the same type
	GetRooms(ctx context.Context, user User, Type ChatType) ([]ChatRoom, error)
	// Get all the Members of a room
	GetMembers(ctx context.Context, Room ChatRoom) ([]ChatRoom, error)
	// Get all the Room that the User is joining
	GetRoomsFromUser(ctx context.Context, user User) ([]ChatRoom, error)

	// Delets the proviede room
	RemoveRoom(ctx context.Context, Room ChatRoom) error
}
