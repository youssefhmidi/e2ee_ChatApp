package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

func newLoginRoute(engine *gin.Engine, lr models.LoginRoute) {
	engine.POST("/login", lr.LoginHandler)
}
