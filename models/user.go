package models

import (
	"context"

	"gorm.io/gorm"
)

// main model
type User struct {
	gorm.Model
	PublicKey      string `gorm:"unique"`
	Name           string
	Email          string `gorm:"unique"`
	Pasword        string
	NickName       string
	Messages       []Message  `gorm:"foreignKey:UserID"`
	ChatRoomsOwned []ChatRoom `gorm:"foreignKey:OwnerID"`
	ChatRooms      []ChatRoom `gorm:"many2many:user_chatroom"`
}

// request and response models
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// interfacess
type UserRepository interface {
	//	Create a user
	CreatUser(ctx context.Context, user User) error

	//	get a user

	GetUserById(ctx context.Context, ID uint) (User, error)
	GetUserByPublicKey(ctx context.Context, publicKey string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	//	fetching all users
	FetchAllUsers(ctx context.Context, limit int) ([]User, error)

	//	Update Information about a User

	//	     // target : a table column like "email, password, ...etc", E.g:
	//	  	UpdateUser(ctx, user, "name", "NewName")
	UpdateUser(ctx context.Context, user User, target string, value interface{}) error
	//   // AppendToUser is like the built in append but with an extra argument wich is the assosiation
	//	 // the assosiation is a field Name in the struct type not in the sql table
	//	 // make sure the in arg is a Slice
	AppendToUser(ctx context.Context, user User, assosiation string, in interface{}) error

	// delete User

	DeleteUser(ctx context.Context, user User) error
}

type UserService interface {
	// Gets the user by the access token provided
	GetUserByToken(ctx context.Context, token string) (User, error)
	// Refresh the acces token and return new access token and another refresh token
	RefreshToken(ctx context.Context, refreshToken string) (AuthResponse, error)
}

type LoginService interface {
	// checks if the email exist
	ValidateEmail(ctx context.Context, email string) bool
	// check if the password passed in are to the provided email
	ValidateUser(ctx context.Context, request LoginRequest) bool
	// Create a respose with two tokens (access and refresh)
	StartJwtSession(user User) AuthResponse
}

type SignupService interface {
	// Checks if the email alreay exist
	IsEmailExist(ctx context.Context, email string) bool
	// Register a User
	RegisterUser(ctx context.Context, request SignupRequest) AuthResponse
}
