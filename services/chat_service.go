package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type ChatSerive struct {
	MessageRepository models.MessageRepository
}

func NewChatService(mr models.MessageRepository) models.ChatService {
	return &ChatSerive{
		MessageRepository: mr,
	}
}

func (cs *ChatSerive) VerifyMessage(ctx context.Context, sender models.User, signedMessage string) error {
	message, IsVerified := VerifySignature(signedMessage, sender.PublicKey)
	if !IsVerified {
		return errors.New(fmt.Sprintf("the signature of %v do not match the signature of user:%v is not verified", message["message"], sender.ID))
	}
	return nil
}

func (cs *ChatSerive) SendMessage(ctx context.Context, message models.Message) error {
	return cs.MessageRepository.CreateMessage(ctx, message)
}
