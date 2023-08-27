package models

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// types

// can either be "group" or "dm"
type ChatType string

type ChatRoom struct {
	gorm.Model
	OwnerID   uint
	Name      string `gorm:"unique"`
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
	// this function should be used in the main function to get all the rooms
	GetRooms() ([]ChatRoom, error)

	// fetching multiple ChatRoom
	GetRoomsFromUser(ctx context.Context, limit int, user User) ([]ChatRoom, error)
	GetOwnedRooms(ctx context.Context, limit int, user User) ([]ChatRoom, error)
	GetRoomsByType(ctx context.Context, Type ChatType, limit int) ([]ChatRoom, error)
	GetMembers(ctx context.Context, room ChatRoom, limit int) ([]User, error)

	// Updating a ChatRoom
	UpdateRoom(ctx context.Context, room ChatRoom, target string, value interface{}) error
	AppendToRoom(ctx context.Context, room ChatRoom, association string, in interface{}) error

	// Delete a ChatRoom
	DeleteRoom(ctx context.Context, room ChatRoom) error
	DeleteFromRoom(ctx context.Context, room ChatRoom, association string, in interface{}) error
}

type ChatRoomService interface {
	// Creates Group and initilize a symetric encryption key
	//
	// ``` the Member parameter is optionnal ```
	//
	// returns a key and a error (which is nil if the operation does correctly)
	CreateGroup(ctx context.Context, Name string, Owner User, Members []User, IsPublic bool) (string, error)
	// Creates a DM with another user.
	//
	// user1 and user2 are the participent in the DM
	//
	// returns an error if not succeded
	CreateDM(ctx context.Context, user1 User, user2 User) error

	// Add a member to the provided room
	// returns an error if failed
	AddMember(ctx context.Context, Room ChatRoom, user User) error
	// Removes a member from the provided room
	// returns an error if failed
	RemoveMember(ctx context.Context, Room ChatRoom, user User) error

	// Gets all ChatRooms with the same type
	GetRooms(ctx context.Context, Type ChatType) ([]ChatRoom, error)
	// Get all the Members of a room
	GetMembers(ctx context.Context, Room ChatRoom) ([]User, error)
	// Get all the Room that the User is joining
	GetRoomsFromUser(ctx context.Context, user User) ([]ChatRoom, error)

	// Delete the proviede room
	RemoveRoom(ctx context.Context, Room ChatRoom) error

	// Return a chatroom by its passed 'val' parameter
	//
	// val can be a string or a uint, if another type is passed it will return a ErrInvildeParameterType error and an empty chat room
	GetRoomBy(ctx context.Context, val any) (ChatRoom, error)

	// generate a token that expires after 1 hour and should be used to verify
	// users inside the connection
	GenerateSessionToken(user User) (string, error)
}

// Router interface for Handlers
//
// TODO : add a enpoint of radding users and inviting them
//
// the Endpoint will be in a group '/chat/'
type RoomRouter interface {
	// this handler must be used by the relative path '/chatserver/:room_id/join'
	// and a invite key may be nessecairy if the room is not public
	JoinHandler(c *gin.Context)

	// this handler must be used by the relative path '/chat/new'
	// this will create a room and the user should only care about the name and if its public or not
	CreateRoomHandler(c *gin.Context)

	// get a specific data about a room
	// the 'requested' param is what the database will send to the user
	// relative path : '/chat/:room_id'
	GetHandler(c *gin.Context)

	// generates a token for the user to use to connect to the server
	// it has an expiry duration of an hour
	// relative path /chat/session
	GenerateToken(c *gin.Context)
}

// JoinRequest is the provided request the user should give in order to join a room
//
// this is optionnal only if the room is privet and the user is joining for the first time
type JoinRequest struct {
	InviteKey string `json:"invite_key"`
}

type CreateRoomRequest struct {
	Name      string   `json:"name"`
	MembersID []uint   `json:"members_id"`
	IsPublic  bool     `json:"is_public"`
	Type      ChatType `json:"type"`
}
