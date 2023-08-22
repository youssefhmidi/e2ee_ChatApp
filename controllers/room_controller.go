package controllers

import (
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/websocket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type RoomController struct {
	SocketServer websocket.SocketServer
	RoomService  models.ChatRoomService
	ChatService  models.ChatService
}
