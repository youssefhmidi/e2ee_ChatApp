package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type Client struct {
	// Conn is going to be used by the Room to be able to brodcast messages to all members of a room
	Conn websocket.Conn

	// the information about the connected client
	User *models.User

	// chanel for sending strings
	Send chan []byte

	// the room that the client is currently connecting
	Room *Room
}
