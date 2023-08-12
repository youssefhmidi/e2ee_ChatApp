package services

import (
	"context"

	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/auth"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type UserService struct {
	UserRepository models.UserRepository
	JwtService     models.JwtService
}

func NewUserService(ur models.UserRepository, JwtS models.JwtService) models.UserService {
	return &UserService{
		UserRepository: ur,
		JwtService:     JwtS,
	}
}

func (us *UserService) GetUserByToken(ctx context.Context, token string) (models.User, error) {
	Id, err := auth.GetIdFromToken(token, us.JwtService.GetSecret("access"))
	if err != nil {
		return models.User{}, err
	}
	user, err := us.UserRepository.GetUserById(ctx, Id)
	return user, err
}
func (us *UserService) RefreshToken(ctx context.Context, refreshToken string) models.AuthResponse {

}
