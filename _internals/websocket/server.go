package websocket

import (
	"errors"

	"github.com/gorilla/websocket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

var (
	DefaultUpgrader = websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
	}
	ErrRoomNotFound = errors.New("couldn't find the room in socket server Rooms, consider talking to")
)

// the SocketServer will manage every room in the server and will loop over every one and start one by one
type SocketServer struct {
	// Rooms are list of rooms that will be started seperatlly in a goroutine
	Rooms []Room
}

// Getting the Room by its Room field,
// which is just the 'models.ChatRoom' field
func (ss *SocketServer) GetRoom(ChatRoom models.ChatRoom) (Room, error) {
	for _, room := range ss.Rooms {
		if room.Room.ID == ChatRoom.ID {
			return room, nil
		}
	}
	return Room{}, ErrRoomNotFound
}

// a Service that the room controller will use to make rooms and to open
// websocket connection and start a rooms
type WebSocketService interface {
	// checks the user's rooms and returns true if the access is accepted or false
	// if the user is not a part of that room
	VerifyAccess(usr models.User, room models.ChatRoom) bool

	// add user to the room and start a Client
	JoinRoom(usr models.User, room Room) error
}

func StartRooms(ChatRooms []models.ChatRoom) {
	for _, r := range ChatRooms {
		room := NewRoom(r)
		go room.Run()
	}
}
