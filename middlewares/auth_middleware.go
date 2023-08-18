package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/auth"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

// this function returns a middleware function to use
//
// you can set up which type of verification just by puting in the right secret
func UseTokenVerification(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorisation")

		t, accessToken := auth.GetAuthType(token)
		if !auth.IsBearer(t) {
			c.Abort()
			return
		}
		IsValide, err := auth.ValidateToken(accessToken, secret)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorMessage{ResponseMessage: "can't validate the given token, error : " + err.Error() + ", if the err seems to be reappearing, pls consider talking to an admin"})
			c.Abort()
			return
		}
		if !IsValide {
			c.JSON(http.StatusBadRequest, models.ErrorMessage{ResponseMessage: "unAuthorized jwt Token"})
		}
	}
}
