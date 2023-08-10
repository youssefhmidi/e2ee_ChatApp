package repository

import (
	"context"

	"github.com/youssefhmidi/E2E_encryptedConnection/database"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type UserRepository struct {
	Db database.SqliteDatabase
}

func NewUserRepository(db database.SqliteDatabase) models.UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (ur *UserRepository) CreatUser(ctx context.Context, user models.User) error {
	return ur.Db.Insert(ctx, &user)
}
func (ur *UserRepository) GetUserById(ctx context.Context, ID uint) (models.User, error) {
	res, err := ur.Db.GetModelById(ctx, &models.User{}, ID)
	return res.(models.User), err
}
func (ur *UserRepository) GetUserByPublicKey(ctx context.Context, publicKey string) (models.User, error) {
	res, err := ur.Db.GetModelWhere(ctx, &models.User{}, "public_key", publicKey)
	return res.(models.User), err
}
func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	res, err := ur.Db.GetModelWhere(ctx, &models.User{}, "email", email)
	return res.(models.User), err
}
func (ur *UserRepository) FetchAllUsers(ctx context.Context, limit int) ([]models.User, error) {
	res, err := ur.Db.GetAll(ctx, limit, &[]models.User{})
	return res.([]models.User), err
}
func (ur *UserRepository) UpdateUser(ctx context.Context, user models.User, target string, value interface{}) error {
	return ur.Db.UpdateModel(ctx, &user, target, value)
}
func (ur *UserRepository) AppendToUser(ctx context.Context, user models.User, association string, in interface{}) error {
	return ur.Db.AppendTo(ctx, &user, association, &in)
}
func (ur *UserRepository) DeleteUser(ctx context.Context, user models.User) error {
	return ur.Db.DeleteModel(ctx, &user)
}
