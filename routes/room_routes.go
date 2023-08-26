package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

func newRoomRoutes(roomGroup *gin.RouterGroup, rr models.RoomRouter) {
	roomGroup.GET("/:room_id", rr.GetHandler)
	roomGroup.POST("/new", rr.CreateRoomHandler)
	roomGroup.GET("/:room_id/join", rr.JoinHandler)
}
