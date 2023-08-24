package socket

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

// types that this package and main packages will use
type (
	ClientMessage struct {
		// encrypted message is the input that the client send to the server
		EncryptedMessage string `json:"encrypted_message"`
		// the sender ID this will be used to create a message object in the db
		SenderID uint `json:"sender"`
		// this field also going to be used to store the message in the db
		ChatRoomID uint `json:"room_id"`
	}
	// an alias so It will be esy to document the type
	ClientMessageCh chan ClientMessage

	// a Store is a list of channels so the server will be able to seperatlly add messages to the server
	Store map[*Room]ClientMessageCh

	// StorageFunc is a function that take care of storing the Recived data from the messages channel that is maped by its Room
	// it returns an error if it can't store
	StorageFunc func(ClientMessageCh) error
)

// methods for the Store type alias
func (s Store) Store(room *Room, storageFunc StorageFunc) {
	for {
		if err := storageFunc(s[room]); err != nil {
			log.Fatal(err, "IT WAS HERE FROM THE BEGENING")
			break
		}
	}
}

var (
	// a already pre made Ugrader that has a
	// WriteBufferSize and ReadBufferSize of 1024
	DefaultUpgrader = websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ErrRoomNotFound = errors.New("couldn't find the room in socket server Rooms, consider talking to")
	ErrStoringFaild = errors.New("can't store the message to the database")
)

// the SocketServer will manage every room in the server and will loop over every one and start one by one
type SocketServer struct {
	// Rooms are list of rooms that will be started seperatlly in a goroutine
	Rooms []*Room

	// a Store that will storing messages to the db
	LocalStore Store

	// Storing function that will be executed by the store
	StorageFunc
}

// Getting the Room by its Room field,
// which is just the 'models.ChatRoom' field
func (ss *SocketServer) GetRoom(ChatRoom models.ChatRoom) (*Room, error) {
	for _, room := range ss.Rooms {
		if room.ChatRoom.ID == ChatRoom.ID {
			return room, nil
		}
	}
	return &Room{}, ErrRoomNotFound
}

// initilize a SocketServer and populates it with Rooms
func (ss *SocketServer) InitAndRun(DBChatRooms []models.ChatRoom) {
	// Add registered Rooms to the SocketServer
	for _, r := range DBChatRooms {
		ss.Rooms = append(ss.Rooms, NewRoom(r))
	}
	// run the main loop for the rooms
	startServer(ss.Rooms, ss)
}

// start the server and hosts all the rooms
func startServer(ChatRooms []*Room, s *SocketServer) {
	s.LocalStore = make(Store)
	for _, r := range ChatRooms {
		s.LocalStore[r] = make(ClientMessageCh, 30)

		go r.Run(s.LocalStore)
		go s.LocalStore.Store(r, s.StorageFunc)
	}
}

// a Service that the room controller will use to make rooms and to open
// websocket connection and start a rooms
type WebSocketService interface {
	// checks the user's rooms and returns true if the access is accepted or false
	// if the user is not a part of that room
	VerifyAccess(ctx context.Context, usr models.User, room models.ChatRoom) bool

	// Creat a client with the ptovided websocket connection ,user and room
	CreateClient(ctx context.Context, ws *websocket.Conn, usr models.User, room Room) (*Client, error)

	// return a storagefunc that will store all the messages
	StoreMsgsToDatabase(ctx context.Context) StorageFunc
}
