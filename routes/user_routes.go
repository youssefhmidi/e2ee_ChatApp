package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/middlewares"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

// creates a /refresh endpoint
func newRefreshRoute(engine *gin.Engine, ur models.UserRoute) {
	engine.GET("/refresh", middlewares.UseTokenVerification("refresh", "refresh"), ur.RefreshTokenHandler)
}

// creates a /users/@me
func newUserRoute(userGroup *gin.RouterGroup, ur models.UserRoute) {
	userGroup.GET("/@me", ur.UserHandler)
}
