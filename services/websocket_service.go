package services

import "github.com/youssefhmidi/E2E_encryptedConnection/models"

type WebsockeService struct {
	ChatService models.ChatService
	RoomService models.ChatRoomService
}
