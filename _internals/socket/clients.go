package socket

import (
	"log"

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
	Send ClientMessageCh

	// the room that the client is currently connecting
	Room *Room
}

// initilize a client
func NewClient(user *models.User, conn *websocket.Conn, room *Room) *Client {
	return &Client{
		User: user,
		Conn: conn,
		Send: make(ClientMessageCh),
		Room: room,
	}
}

// ReadIn will read every input that the user will provide and brodcast it to the whole room
func (c *Client) ReadIn() {
	// closing the connection as soon as the loop breaks
	defer func() {
		c.Conn.Close()
		c.Room.Leave <- c
	}()

	// loop for brodcasting every message to the room
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
				log.Println("error :", err)
			}
			break
		}
		log.Println("read message")
		Message := ClientMessage{
			EncryptedMessage: string(msg),
			SenderID:         c.User.ID,
			ChatRoomID:       c.Room.ChatRoom.ID,
		}
		c.Room.Brodcast <- Message
	}
}

// WriteOut will send to the client every message that has been brodcasted to the Room the Client is curently in
func (c *Client) WriteOut() {

	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:

			// if somthing happened break the loop
			if !ok {
				break
			}

			// write to the websocket connection the message
			if err := c.Conn.WriteJSON(message); err != nil {
				log.Println("error :", err)
			}
		default:
			continue
		}
	}
}
