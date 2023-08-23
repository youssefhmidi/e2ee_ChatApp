package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/socket"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
	"github.com/youssefhmidi/E2E_encryptedConnection/services"
)

type RoomController struct {
	Env              *bootstraps.Env
	SocketServer     *socket.SocketServer
	WebsocketService socket.WebSocketService
	ChatRoomService  models.ChatRoomService
	GroupChatService services.GroupChatEncryption
}

// todo : finish code here

func NewRoomController(ss *socket.SocketServer, wss socket.WebSocketService, crs models.ChatRoomService, gcs services.GroupChatEncryption, env *bootstraps.Env) models.RoomRouter {
	return &RoomController{
		SocketServer:     ss,
		WebsocketService: wss,
		ChatRoomService:  crs,
		GroupChatService: gcs,
		Env:              env,
	}
}

func (rc *RoomController) JoinHandler(c *gin.Context) {

}

func (rc *RoomController) CreateRoomHandler(c *gin.Context) {

}

func (rc *RoomController) AddMemberHandler(c *gin.Context) {

}

func (rc *RoomController) RemoveMemberHandler(c *gin.Context) {

}

func (rc *RoomController) GetHandler(c *gin.Context) {
	// Should return is to know if the function should send the Output map
	ShouldReturn := false
	// the key of the Output map is what the user asked for and the value will be the result
	// similar to GraphQl
	Output := make(map[string]interface{})

	// converting the param from string to int to uint
	RoomIdStr := c.Param("room_id")
	ID, err := strconv.ParseUint(RoomIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}
	RoomID := uint(ID)

	// Getting the room by its id
	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()
	room, err := rc.ChatRoomService.GetRoomBy(ctx, RoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	// checking if the user want room's members
	_, IsMember := c.GetQuery("members")
	if IsMember {
		ShouldReturn = true
		ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
		defer cancel()
		Members, err := rc.ChatRoomService.GetMembers(ctx, room)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
			return
		}
		Output["members"] = Members
	}

	// checks if there the Owner
	_, IsOwner := c.GetQuery("owner_id")
	if IsOwner {
		ShouldReturn = true
		Output["owner_id"] = room.OwnerID
	}
	if ShouldReturn {
		c.JSON(http.StatusOK, Output)
	}

	c.JSON(http.StatusOK, room)
}
