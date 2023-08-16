package services

import "github.com/youssefhmidi/E2E_encryptedConnection/models"

type RoomService struct {
	UserRepository    models.UserRepository
	EncryptionService GroupChatEncryption
	RoomRepository    models.ChatRoomRepository
}

func NewRoomService(ur models.UserRepository, es GroupChatEncryption, Rr models.ChatRoomRepository) {

}
