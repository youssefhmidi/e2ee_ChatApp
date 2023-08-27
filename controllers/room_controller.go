package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/socket"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type RoomController struct {
	Env              *bootstraps.Env
	SocketServer     *socket.SocketServer
	WebsocketService socket.WebSocketService
	ChatRoomService  models.ChatRoomService
	UserService      models.UserService
}

func NewRoomController(ss *socket.SocketServer, wss socket.WebSocketService, crs models.ChatRoomService, us models.UserService, env *bootstraps.Env) models.RoomRouter {
	return &RoomController{
		SocketServer:     ss,
		WebsocketService: wss,
		UserService:      us,
		ChatRoomService:  crs,
		Env:              env,
	}
}

// needs a rewrite
func (rc *RoomController) JoinHandler(c *gin.Context) {
	// Upgrade turn the request into a websocket connection. and registrating the room
	ws, err := socket.DefaultUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}
	// getting the user access token
	accessToken := c.MustGet("access_token")

	// getting the athor of the request
	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()
	user, err := rc.UserService.GetUserByToken(ctx, accessToken.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	// getting the specified room from the request
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()
	RoomID, err := toUint(c.Param("room_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}
	room, err := rc.ChatRoomService.GetRoomBy(ctx, RoomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			ResponseMessage: `Can not find the room with the provided id, make sure that th id is correct, 
			no typos and if the err seems to occur again contact an admin, Error :` + err.Error(),
		})
		return
	}

	// Checking if the user have access to the room and verifing the user invite key
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()
	if !rc.WebsocketService.VerifyAccess(ctx, user, room) {
		// TODO : add a invite key logic
		c.JSON(http.StatusBadRequest, "man! I don't know either wtf to do a this point")
	}

	room_Conn, err := rc.SocketServer.GetRoom(room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	// Starting a client and adding the user to the room
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()
	client, err := rc.WebsocketService.CreateClient(ctx, ws, user, *room_Conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}
	room_Conn.Join <- client
	go client.ReadIn()
	go client.WriteOut()
}

func (rc *RoomController) CreateRoomHandler(c *gin.Context) {
	// binding the request to the req variable
	var req models.CreateRoomRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	// getting the user access token
	accessToken := c.MustGet("access_token")

	// getting the athor of the request
	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()
	user, err := rc.UserService.GetUserByToken(ctx, accessToken.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	// getting the users by their id
	members := []models.User{}
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()
	for _, IDs := range req.MembersID {
		usr, err := rc.UserService.GetUserById(ctx, IDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
			return
		}
		members = append(members, usr)
	}

	// creating the room
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()

	if req.Type == "group" {
		key, err := rc.ChatRoomService.CreateGroup(ctx, req.Name, user, members, req.IsPublic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, models.SuccesMessage{ResponseMessage: "Chat Group created ,EncryptionKey :" + key})
		return
	}

	if err := rc.ChatRoomService.CreateDM(ctx, user, members[0]); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: " cant create room" + err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, models.SuccesMessage{ResponseMessage: "DM created"})

	// adding the room to the SocketServer and runing them seperatly
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(rc.Env.ContextTimeout))
	defer cancel()
	room, err := rc.ChatRoomService.GetRoomBy(ctx, req.Name)
	if err != nil {
		log.Fatal(err)
	}
	rc.SocketServer.RunAndRegisterRoom(room)
}

func (rc *RoomController) GetHandler(c *gin.Context) {
	// Should return is to know if the function should send the Output map
	ShouldReturn := false
	// the key of the Output map is what the user asked for and the value will be the result
	// similar to GraphQl
	Output := make(map[string]interface{})

	// converting the param from string to int to uint
	RoomIdStr := c.Param("room_id")
	RoomID, err := toUint(RoomIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

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
		return
	}

	c.JSON(http.StatusOK, room)
}

func toUint(paramReq string) (uint, error) {
	ID, err := strconv.ParseUint(paramReq, 10, 64)
	RoomID := uint(ID)
	return RoomID, err
}
