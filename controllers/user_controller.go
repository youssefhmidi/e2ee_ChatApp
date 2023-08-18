package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type UserController struct {
	Env         *bootstraps.Env
	UserService models.UserService
}

func NewUserController(env *bootstraps.Env, us models.UserService) models.UserRoute {
	return &UserController{
		Env:         env,
		UserService: us,
	}
}

// this handler must be used by the /users/@me route
func (uc *UserController) UserHandler(c *gin.Context) {
	access := c.MustGet("access_token")
	token := access.(string)

	// getting the user
	ctx, cancel := context.WithTimeout(c, time.Duration(uc.Env.ContextTimeout)*time.Second)
	defer cancel()
	user, err := uc.UserService.GetUserByToken(ctx, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}
	// creating and returning the data that the user will be caring of
	userResp := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		NickName:  user.NickName,
		CreatedAt: user.CreatedAt,
		Email:     user.Email,
		PublicKey: user.PublicKey,
	}
	c.JSON(http.StatusOK, userResp)
}

// this should be used by the /refresh/
func (uc *UserController) RefreshTokenHandler(c *gin.Context) {
	refresh := c.MustGet("refresh_token")

	// refreshing the token
	// if you're wondering where is the secret, the sicret is actually passed in when useing the NewUserService function
	ctx, cancel := context.WithTimeout(c, time.Duration(uc.Env.ContextTimeout)*time.Second)
	defer cancel()
	resp, err := uc.UserService.RefreshToken(ctx, refresh.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
