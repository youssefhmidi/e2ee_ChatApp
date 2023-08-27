package repository

import (
	"context"

	"github.com/youssefhmidi/E2E_encryptedConnection/database"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type ChatRoomRepository struct {
	Db database.SqliteDatabase
}

func NewChatRepository(db database.SqliteDatabase) models.ChatRoomRepository {
	return &ChatRoomRepository{
		Db: db,
	}
}

func (cr *ChatRoomRepository) CreateChatRoom(ctx context.Context, room models.ChatRoom) error {
	return cr.Db.Insert(ctx, &room)
}

func (cr *ChatRoomRepository) GetRoomByID(ctx context.Context, ID uint) (models.ChatRoom, error) {
	res, err := cr.Db.GetModelById(ctx, &models.ChatRoom{}, ID)
	return *res.(*models.ChatRoom), err
}

func (cr *ChatRoomRepository) GetRoomByName(ctx context.Context, Name string) (models.ChatRoom, error) {
	res, err := cr.Db.GetModelWhere(ctx, &models.ChatRoom{}, "name", Name)
	return *res.(*models.ChatRoom), err
}
func (cr *ChatRoomRepository) GetRoomsByType(ctx context.Context, Type models.ChatType, limit int) ([]models.ChatRoom, error) {
	res, err := cr.Db.GetAllWhere(ctx, limit, &[]models.ChatRoom{}, "type", Type)
	return *res.(*[]models.ChatRoom), err
}

func (cr *ChatRoomRepository) GetRoomsFromUser(ctx context.Context, limit int, user models.User) ([]models.ChatRoom, error) {
	res, err := cr.Db.GetAllWithAssociation(ctx, limit, &user, &[]models.ChatRoom{}, "ChatRooms")
	return *res.(*[]models.ChatRoom), err
}

func (cr *ChatRoomRepository) GetOwnedRooms(ctx context.Context, limit int, user models.User) ([]models.ChatRoom, error) {
	res, err := cr.Db.GetAllWhere(ctx, limit, &[]models.ChatRoom{}, "owner_id", user.ID)
	return *res.(*[]models.ChatRoom), err
}

func (cr *ChatRoomRepository) GetMembers(ctx context.Context, room models.ChatRoom, limit int) ([]models.User, error) {
	res, err := cr.Db.GetAllWithAssociation(ctx, limit, &room, &[]models.User{}, "Members")
	return *res.(*[]models.User), err
}

func (cr *ChatRoomRepository) UpdateRoom(ctx context.Context, room models.ChatRoom, target string, value interface{}) error {
	return cr.Db.UpdateModel(ctx, &room, target, value)
}

func (cr *ChatRoomRepository) AppendToRoom(ctx context.Context, room models.ChatRoom, association string, in interface{}) error {
	return cr.Db.AppendTo(ctx, &room, association, &in)
}

func (cr *ChatRoomRepository) DeleteRoom(ctx context.Context, room models.ChatRoom) error {
	return cr.Db.DeleteModel(ctx, &room)
}

func (cr *ChatRoomRepository) DeleteFromRoom(ctx context.Context, room models.ChatRoom, association string, in interface{}) error {
	return cr.Db.DeleteAssociation(ctx, &room, association, &in)
}

func (cr *ChatRoomRepository) GetRooms() ([]models.ChatRoom, error) {
	res, err := cr.Db.GetAll(context.Background(), 40, &[]models.ChatRoom{})
	return *res.(*[]models.ChatRoom), err
}
