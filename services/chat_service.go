package services

import "github.com/youssefhmidi/E2E_encryptedConnection/models"

type ChatSerive struct {
	UserRepository    models.UserRepository
	MessageRepository models.MessageRepository
	ChatRoomService   RoomService
}
