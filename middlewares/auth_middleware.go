package middlewares

import (
	"log"
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
		token := c.GetHeader("Authorization")

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

// a rewrite for makin a websocket auth
func UseWebsocketAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// getting the token from the query and checking if its provided
		token, IsProvided := c.GetQuery("token")
		if !IsProvided {
			c.Abort()
			return
		}

		// validating if it is valid or not otherwise it abort the connection
		IsValide, err := auth.ValidateToken(token, secret)
		if !IsValide || err != nil {
			log.Fatal("cannot validat token , error :", err)
			c.Abort()
			return
		}

		userId, err := auth.GetIdFromToken(token, secret)
		if err != nil {
			log.Fatal("Got Error :", err)
		}
		// setting an access token
		c.Set("user_id", userId)
		c.Next()
	}
}
