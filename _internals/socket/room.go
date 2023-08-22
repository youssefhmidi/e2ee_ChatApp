package socket

import (
	"github.com/gorilla/websocket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type Room struct {
	// Room field is for verifying users and for extra functionalities
	ChatRoom models.ChatRoom

	// clients who are currently connected to the room
	Clients map[*Client]*websocket.Conn

	// Brodcast is a channel that is used to brodcst any message to every other connection in the room
	Brodcast ClientMessageCh

	// channel for joining a room
	Join chan *Client

	// channel for leaving
	Leave chan *Client
}

func NewRoom(r models.ChatRoom) *Room {
	return &Room{
		ChatRoom: r,
		Clients:  make(map[*Client]*websocket.Conn),
		Brodcast: make(ClientMessageCh),
		Join:     make(chan *Client),
		Leave:    make(chan *Client),
	}
}

// TODOS : Need a re-write so it will be posible to store messages
func (r *Room) Run(store Store) {
	ClientMsgch := store[r]
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = client.Conn
		case client := <-r.Leave:
			close(client.Send)
			delete(r.Clients, client)
		case message := <-r.Brodcast:
			ClientMsgch <- message

			for client := range r.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(r.Clients, client)
				}
			}
		}
	}
}
