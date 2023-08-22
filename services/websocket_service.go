package services

import (
	"context"
	"log"

	"github.com/gorilla/websocket"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/socket"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type WebsockeService struct {
	ChatService models.ChatService
	RoomService models.ChatRoomService
}

// todo implement the  WebsocketService interface
func NewWebsocketService(cs models.ChatService, rs models.ChatRoomService) socket.WebSocketService {
	return &WebsockeService{
		ChatService: cs,
		RoomService: rs,
	}
}

func (wss *WebsockeService) VerifyAccess(ctx context.Context, usr models.User, room models.ChatRoom) bool {
	// checking if the room is public or not
	if room.IsPublic {
		return true
	}

	// getting the list of all members in the room
	members, err := wss.RoomService.GetMembers(ctx, room)
	if err != nil {
		log.Fatal(err)
	}

	// checking if the user is part of the chatroom
	for _, user := range members {
		if user.ID == usr.ID {
			return true
		}
	}

	// return false if none of the above condition are met
	return false

}

func (wss *WebsockeService) CreateClient(ctx context.Context, ws *websocket.Conn, usr models.User, room socket.Room) (*socket.Client, error) {
	// adding the user to the members list
	err := wss.RoomService.AddMember(ctx, room.ChatRoom, usr)

	// creating a client object the client and joing
	client := socket.NewClient(&usr, ws, &room)

	// handling 'err'
	if err == ErrAlreadyInRoom {
		return client, nil
	}

	// returning the client object, and a possible err
	return client, err
}

func (wss *WebsockeService) StoreMsgsToDatabase(ctx context.Context) socket.StorageFunc {
	// this function is meant to be used by the server in the bootstrap directory where
	// the server will be initilized
	return func(cmc socket.ClientMessageCh) error {
		message := <-cmc

		messageDB := models.Message{
			ChatRoomID: message.ChatRoomID,
			UserID:     message.SenderID,
			Message:    message.EncryptedMessage,
		}
		return wss.ChatService.SendMessage(ctx, messageDB)
	}
}
