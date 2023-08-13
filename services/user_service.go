package services

import (
	"context"
	"errors"

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
	Id, err := auth.GetIdFromToken(token, us.JwtService.GetSecret(KeyForAccess))
	if err != nil {
		return models.User{}, err
	}
	// gets the User by its ID and return the User and an error if it has one
	user, err := us.UserRepository.GetUserById(ctx, Id)
	return user, err
}
func (us *UserService) RefreshToken(ctx context.Context, RefreshToken string) (models.AuthResponse, error) {
	// Validate the token and return an empty struct object and an error if it catch some errors
	IsValid, err := auth.ValidateToken(RefreshToken, us.JwtService.GetSecret(KeyForRefresh))
	if err != nil {
		return models.AuthResponse{}, err
	}

	// checks if its valide if not return an empty struct object and an error
	if !IsValid {
		return models.AuthResponse{}, errors.New("invalid Token")
	}

	// gets the ID from the refreshToken argument otherwise returns an empty struct object and an error
	ID, err := auth.GetIdFromToken(RefreshToken, us.JwtService.GetSecret(KeyForRefresh))
	if err != nil {
		return models.AuthResponse{}, err
	}

	// getting a User struct object by its ID otherwise returns an empty struct object and an error
	usr, err := us.UserRepository.GetUserById(ctx, ID)
	if err != nil {
		return models.AuthResponse{}, err
	}

	// Generating the access token otherwise returns an empty struct object and an error
	accessToken, err := auth.CreateAccessToken(usr, us.JwtService.GetSecret(KeyForAccess), us.JwtService.GetExpiryTime(KeyForAccess))
	if err != nil {
		return models.AuthResponse{}, err
	}

	// Generating the Refresh token otherwise returns an empty struct object and an error
	refreshToken, err := auth.CreateRefreshToken(usr, us.JwtService.GetSecret(KeyForRefresh), us.JwtService.GetExpiryTime(KeyForRefresh))
	if err != nil {
		return models.AuthResponse{}, err
	}

	// returns an Authorization response and a 'nil' error
	return models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
