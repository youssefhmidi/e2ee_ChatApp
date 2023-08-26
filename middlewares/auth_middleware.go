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
// and the usage argument is meant for making sure what kind of response you want to pass the the pending handler
//
// i.e for access_token the usage will be access and the passed key/val pair is "access_token"
func UseTokenVerification(secret string, usage string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// getting the Authorisation value from the header
		token := c.GetHeader("Authorisation")

		// gets the auth type to validate if the Header request is in this shape
		// 'Bearer <Token>'
		t, accessToken := auth.GetAuthType(token)
		if !auth.IsBearer(t) {
			c.Abort()
			return
		}

		// validate that the token is indeed signed with the access token secret
		IsValide, err := auth.ValidateToken(accessToken, secret)
		// handle the err if there are an error
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorMessage{ResponseMessage: "can't validate the given token, error : " + err.Error() + ", if the err seems to be reappearing, pls consider talking to an admin"})
			c.Abort()
			return
		}
		// checks if its a valide token
		if !IsValide {
			c.JSON(http.StatusBadRequest, models.ErrorMessage{ResponseMessage: "unAuthorized jwt Token"})
		}

		// store a key/value pair with the structure below
		// access_token = <accessToken>
		c.Set(usage+"_token", accessToken)
		c.Next()
	}
}
