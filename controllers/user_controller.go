package controllers

import (
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

func (uc *UserController) UserHandler(c *gin.Context) {

}
