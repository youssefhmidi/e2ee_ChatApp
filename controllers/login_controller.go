package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type LoginController struct {
	Env          *bootstraps.Env
	LoginService models.LoginService
}

func NewLoginController(env *bootstraps.Env, ls models.LoginService) models.LoginRoute {
	return &LoginController{
		Env:          env,
		LoginService: ls,
	}
}

func (lc *LoginController) LoginHandler(c *gin.Context) {
	// binding the request body to the req variable and checking if the structure is correct
	var req models.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	// checks if the email does exist
	ctx, cancel := context.WithTimeout(c, time.Duration(lc.Env.ContextTimeout)*time.Second)
	defer cancel()
	if !lc.LoginService.ValidateEmail(ctx, req.Email) {
		c.JSON(http.StatusConflict, models.ErrorMessage{ResponseMessage: "email do not exist"})
		return
	}

	// validating the password for the provided email
	ctx, cancel = context.WithTimeout(c, time.Duration(lc.Env.ContextTimeout)*time.Second)
	defer cancel()
	if !lc.LoginService.ValidateUser(ctx, req) {
		c.JSON(http.StatusConflict, models.ErrorMessage{ResponseMessage: "wrong password"})
		return
	}

	// getting the user by its email (email is a unique row) and creating a access.refresh key pair
	ctx, cancel = context.WithTimeout(c, time.Duration(lc.Env.ContextTimeout)*time.Second)
	defer cancel()
	user, err := lc.LoginService.GetUserByEmail(ctx, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}
	response := lc.LoginService.StartJwtSession(user)

	// returning a access/refresh jwt key pairs
	c.JSON(http.StatusOK, response)
}
