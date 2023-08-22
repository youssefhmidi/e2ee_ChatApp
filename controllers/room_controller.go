package controllers

import (
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/websocket"
)

type RoomController struct {
	SocketServer     websocket.SocketServer
	WebsocketService websocket.WebSocketService
}
