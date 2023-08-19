package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/models"
)

type SignupController struct {
	Env           *bootstraps.Env
	SignupService models.SignupService
}

func NewSignupController(env *bootstraps.Env, ss models.SignupService) models.SignupRoute {
	return &SignupController{
		Env:           env,
		SignupService: ss,
	}
}

func (sc *SignupController) SignupHandler(c *gin.Context) {
	// bnding the information from the request
	var req models.SignupRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	// verifing the user
	ctx, cancel := context.WithTimeout(c, time.Duration(sc.Env.ContextTimeout)*time.Second)
	defer cancel()
	if ok := sc.SignupService.IsEmailExist(ctx, req.Email); !ok {
		c.JSON(http.StatusMethodNotAllowed, models.ErrorMessage{ResponseMessage: "email already exist"})
		return
	}

	// registering the user
	ctx, cancel = context.WithTimeout(c, time.Duration(sc.Env.ContextTimeout)*time.Second)
	defer cancel()
	response, err := sc.SignupService.RegisterUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{ResponseMessage: err.Error()})
		return
	}

	// returning an access/refresh key pair as a succes response
	c.JSON(http.StatusOK, response)
}
