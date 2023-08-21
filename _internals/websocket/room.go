package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type Room struct {
	// Room field is for verifying users and for extra functionalities
	Room models.ChatRoom

	// clients who are currently connected to the room
	Clients map[*Client]*websocket.Conn

	// Brodcast is a channel that is used to brodcst any message to every other connection in the room
	Brodcast chan []byte

	// channel for joining a room
	Join chan *Client

	// channel for leaving
	Leave chan *Client
}

func NewRoom(r models.ChatRoom) *Room {
	return &Room{
		Room:     r,
		Clients:  make(map[*Client]*websocket.Conn),
		Brodcast: make(chan []byte),
		Join:     make(chan *Client),
		Leave:    make(chan *Client),
	}
}

func (r *Room) Run() {

}
