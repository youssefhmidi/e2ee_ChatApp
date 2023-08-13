package services

import (
	"context"
	"log"

	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/auth"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepository models.UserRepository
	JwtService     models.JwtService
}

func NewLoginService(ur models.UserRepository, JwtS models.JwtService) models.LoginService {
	return &LoginService{
		UserRepository: ur,
		JwtService:     JwtS,
	}
}

// return false if the Email do not exist
func (ls *LoginService) ValidateEmail(ctx context.Context, email string) bool {
	if _, err := ls.UserRepository.GetUserByEmail(ctx, email); err == nil {
		return false
	}
	return true
}

func (ls *LoginService) ValidateUser(ctx context.Context, request models.LoginRequest) bool {
	user, err := ls.UserRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return false
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Pasword), []byte(request.Password)); err != nil {
		return false
	}
	return true
}

func (ls *LoginService) StartJwtSession(user models.User) models.AuthResponse {
	Accesstoken, err := auth.CreateAccessToken(user, ls.JwtService.GetSecret(KeyForAccess), ls.JwtService.GetExpiryTime(KeyForAccess))
	if err != nil {
		log.Fatal(err)
		return models.AuthResponse{}
	}
	Refreshtoken, err := auth.CreateRefreshToken(user, ls.JwtService.GetSecret(KeyForRefresh), ls.JwtService.GetExpiryTime(KeyForRefresh))
	if err != nil {
		log.Fatal(err)
		return models.AuthResponse{}
	}

	return models.AuthResponse{
		AccessToken:  Accesstoken,
		RefreshToken: Refreshtoken,
	}

}
