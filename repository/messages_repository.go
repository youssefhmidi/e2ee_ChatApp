package repository

import (
	"context"

	"github.com/youssefhmidi/E2E_encryptedConnection/database"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type MessageRepository struct {
	Db database.SqliteDatabase
}

func NewMessageRepository(db database.SqliteDatabase) models.MessageRepository {
	return &MessageRepository{
		Db: db,
	}
}

func (mr *MessageRepository) CreateMessage(ctx context.Context, message models.Message) error {
	return mr.Db.Insert(ctx, &message)
}

func (mr *MessageRepository) GetMessageByID(ctx context.Context, ID uint) (models.Message, error) {
	res, err := mr.Db.GetModelById(ctx, &models.Message{}, ID)
	return res.(models.Message), err
}

func (mr *MessageRepository) GetMsgsFromUser(ctx context.Context, limit int, user models.User) ([]models.Message, error) {
	res, err := mr.Db.GetAllWhere(ctx, limit, &[]models.Message{}, "user_id", user.ID)
	return res.([]models.Message), err
}
func (mr *MessageRepository) GetMsgsFromRoom(ctx context.Context, limit int, chatRoom models.ChatRoom) ([]models.Message, error) {
	res, err := mr.Db.GetAllWhere(ctx, limit, &[]models.Message{}, "chat_room_id", chatRoom.ID)
	return res.([]models.Message), err
}
func (mr *MessageRepository) UpdateMessage(ctx context.Context, message models.Message, target string, value interface{}) error {
	return mr.Db.UpdateModel(ctx, &message, target, value)
}

func (mr *MessageRepository) DeleteMessage(ctx context.Context, message models.Message) error {
	return mr.Db.DeleteModel(ctx, &message)
}
