package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

func newSignupRoute(engine *gin.Engine, sr models.SignupRoute) {
	engine.POST("/signup", sr.SignupHandler)
}
