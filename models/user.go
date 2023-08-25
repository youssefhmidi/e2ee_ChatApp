package models

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// main model
type User struct {
	gorm.Model
	PublicKey      string `gorm:"unique"`
	Name           string `gorm:"index:idx_name"`
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

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	NickName  string    `json:"nick_name"`
	PublicKey string    `json:"public_key"`
}

type SignupRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	PublicKey string `json:"public_key"`
	Password  string `json:"password"`
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
	// Gets the user by its ID
	GetUserById(ctx context.Context, ID uint) (User, error)
	// Refresh the acces token and return new access token and another refresh token
	RefreshToken(ctx context.Context, refreshToken string) (AuthResponse, error)
	// TODO : Add a Logout function
}

type LoginService interface {
	// checks if the email exist
	ValidateEmail(ctx context.Context, email string) bool
	// check if the password passed in are to the provided email
	ValidateUser(ctx context.Context, request LoginRequest) bool
	// Create a respose with two tokens (access and refresh)
	StartJwtSession(user User) AuthResponse
	// Getting a user by its
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type SignupService interface {
	// Checks if the email alreay exist
	IsEmailExist(ctx context.Context, email string) bool
	// Register a User
	RegisterUser(ctx context.Context, request SignupRequest) (AuthResponse, error)
}

// User interaction route that the API will provide access to
//
// Endpoint : /users/@me
type UserRoute interface {
	// handle Users interaction and return some infotmation about a user
	UserHandler(c *gin.Context)

	// handle a refresh token request
	//
	// Note that this handler is mostly used by a front end dev or an api dev so the user wont be worried about their key getting expired
	RefreshTokenHandler(c *gin.Context)
}

type LoginRoute interface {
	// handle incomming login requests and return a access&refresh token to the requestor
	LoginHandler(c *gin.Context)
}

type SignupRoute interface {
	// handle incomming signup requests and return a access&refresh token to the requestor
	SignupHandler(c *gin.Context)
}
