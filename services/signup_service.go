package services

import (
	"context"

	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/auth"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
	"golang.org/x/crypto/bcrypt"
)

type SignupService struct {
	UserRepository models.UserRepository
	JwtService     models.JwtService
}

func NewSignupService(ur models.UserRepository, jwtS models.JwtService) models.SignupService {
	return &SignupService{
		UserRepository: ur,
		JwtService:     jwtS,
	}
}

func (ss *SignupService) IsEmailExist(ctx context.Context, email string) bool {
	if _, err := ss.UserRepository.GetUserByEmail(ctx, email); err == nil {
		return false
	}
	return true
}

func (ss *SignupService) RegisterUser(ctx context.Context, request models.SignupRequest) (models.AuthResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.AuthResponse{}, err
	}
	usr := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Pasword:   string(password),
		PublicKey: request.PublicKey,
	}

	err = ss.UserRepository.CreatUser(ctx, usr)
	if err != nil {
		return models.AuthResponse{}, err
	}

	accesstoken, err := auth.CreateAccessToken(usr, ss.JwtService.GetSecret(KeyForAccess), ss.JwtService.GetExpiryTime(KeyForAccess))
	if err != nil {
		return models.AuthResponse{}, err
	}

	refreshToken, err := auth.CreateRefreshToken(usr, ss.JwtService.GetSecret(KeyForRefresh), ss.JwtService.GetExpiryTime(KeyForRefresh))
	if err != nil {
		return models.AuthResponse{}, err
	}

	return models.AuthResponse{
		AccessToken:  accesstoken,
		RefreshToken: refreshToken,
	}, nil
}
