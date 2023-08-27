package services

import (
	"context"
	"errors"

	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/auth"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/encryption"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

// errors
var (
	ErrAlreadyInRoom        = errors.New("user already a member in the provided room")
	ErrInvildeParameterType = errors.New("invalid type pls check the function documentation and see the right types that the parameter should be")
)

// the service folder is not documented that well, If you want to know the functionnalities of a function you could see the domain folder
// in this case go to chatroom.go check the ChatRoomService interface to see more

type RoomService struct {
	RoomRepository models.ChatRoomRepository
	JwtService     models.JwtService
}

func NewRoomService(Rr models.ChatRoomRepository, Jwts models.JwtService) models.ChatRoomService {
	return &RoomService{
		RoomRepository: Rr,
		JwtService:     Jwts,
	}
}

func (rs *RoomService) CreateGroup(ctx context.Context, Name string, Owner models.User, Members []models.User, IsPublic bool) (string, error) {
	// generates a gcm object with the type cipher.AEAD
	// and a key wish is the public key that the server will use to encrypte and decrypte the messages
	_, key := encryption.CreateSymetricKey()

	// initializing the room object wich contains member (Optionnal)
	room := models.ChatRoom{
		Name:      Name,
		OwnerID:   Owner.ID,
		PublicKey: key,
		Members:   Members,
		// if IsPublic is true then everyone can join the group
		IsPublic: IsPublic,
		Type:     "group",
	}

	// this err can be some errors with the repository or the db itself
	err := rs.RoomRepository.CreateChatRoom(ctx, room)
	if err != nil {
		return "", err
	}

	// key is a string contains the 32 character-long secret key
	return key, nil
}

func (rs *RoomService) CreateDM(ctx context.Context, user1 models.User, user2 models.User) error {
	// same as CreateGroup function but it automatically create the name and sets the IsPublic to false
	room := models.ChatRoom{
		// there is no owner because (if in the future someone 'or me' added an admin role) a dm no one has authority
		// over some one else
		Name:     user1.Name + "--" + user2.Name + "discussion ",
		IsPublic: false,
		Members:  []models.User{user1, user2},
		Type:     "dm",
	}

	// this err can be some errors with the repository or the db itself
	err := rs.RoomRepository.CreateChatRoom(ctx, room)
	return err
}

func (rs *RoomService) AddMember(ctx context.Context, Room models.ChatRoom, user models.User) error {
	// check if the user is already a member
	members, err := rs.GetMembers(ctx, Room)
	if err != nil {
		return err
	}
	for _, member := range members {
		if member.ID == user.ID {
			return ErrAlreadyInRoom
		}
	}

	// the AppendToRoom method takes the ctx (i.e the context of the request), Room (i.e the room which we want to add a user to)
	// and user (i.e the user we want to add) as arguments and returns an err if something faild
	return rs.RoomRepository.AppendToRoom(ctx, Room, "Members", user)
}

func (rs *RoomService) RemoveMember(ctx context.Context, Room models.ChatRoom, user models.User) error {
	// the DeleteFromRoom is similar to AddMember
	// the only diffrence is it delete records rather than appending them to an association (i.e to the list)
	return rs.RoomRepository.DeleteFromRoom(ctx, Room, "Members", user)
}

func (rs *RoomService) GetRooms(ctx context.Context, Type models.ChatType) ([]models.ChatRoom, error) {
	if Type != "dm" && Type != "group" {
		return []models.ChatRoom{}, errors.New("the Type of the room can be only 'dm' or 'group'")
	}

	// returns all the rooms with this type
	rooms, err := rs.RoomRepository.GetRoomsByType(ctx, Type, 20)
	return rooms, err
}

func (rs *RoomService) GetRoomBy(ctx context.Context, val any) (models.ChatRoom, error) {
	// Checks if the val paramter is type uint and return the result, error
	if ID, ok := val.(uint); ok {
		return rs.RoomRepository.GetRoomByID(ctx, ID)
	}

	if Name, ok := val.(string); ok {
		return rs.RoomRepository.GetRoomByName(ctx, Name)
	}

	// returns an empty ChatRoom object and a Error
	return models.ChatRoom{}, ErrInvildeParameterType
}

func (rs *RoomService) GetMembers(ctx context.Context, Room models.ChatRoom) ([]models.User, error) {
	// easy I know because there's no way to do it efficiently tan this
	rooms, err := rs.RoomRepository.GetMembers(ctx, Room, 20)
	return rooms, err
}

func (rs *RoomService) GetRoomsFromUser(ctx context.Context, user models.User) ([]models.ChatRoom, error) {
	// getting the all rooms where the user is part of
	// this include all the dms and groups (use GetRooms to get specific type of room)
	rooms, err := rs.RoomRepository.GetRoomsFromUser(ctx, 20, user)
	return rooms, err
}

func (rs *RoomService) RemoveRoom(ctx context.Context, Room models.ChatRoom) error {
	// well I guess this is easy to understand right?
	return rs.RoomRepository.DeleteRoom(ctx, Room)
}

func (rs *RoomService) GenerateSessionToken(user models.User) (string, error) {
	// creates a jwt token that expires after 1 hour
	return auth.CreateAccessToken(user, rs.JwtService.GetSecret("socket_access"), rs.JwtService.GetExpiryTime("socket_access"))
}
