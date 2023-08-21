package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

// NOTE: a Client can connect to one room at a time
type Client struct {
	// Conn is going to be used by the Room to be able to brodcast messages to all members of a room
	Conn *websocket.Conn

	// the information about the connected client
	User *models.User

	// chanel for sending strings
	Send chan []byte

	// the room that the client is currently connecting
	Room *Room
}

// initilize a client
func NewClient(user *models.User, conn *websocket.Conn, room *Room) *Client {
	return &Client{
		User: user,
		Conn: conn,
		Send: make(chan []byte),
		Room: room,
	}
}

// ReadIn will read every input that the user will provide and brodcast it to the whole room
func (c *Client) ReadIn() {}

// WriteOut will send to the client every message that has been brodcasted to the Room the Client is curently in
func (c *Client) WriteOut() {}
