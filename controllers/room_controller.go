package controllers

import (
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/socket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
	"github.com/youssefhmidi/E2E_encryptedConnection/services"
)

type RoomController struct {
	SocketServer     socket.SocketServer
	WebsocketService socket.WebSocketService
	ChatRoomService  models.ChatRoomService
	GroupChatService services.GroupChatEncryption
}
