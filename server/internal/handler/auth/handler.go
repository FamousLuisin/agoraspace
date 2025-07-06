package auth

import (
	"fmt"
	"net/http"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/gin-gonic/gin"
)

func NewAuthHandler(service AuthService) AuthHandler {
	return &authHandler{
		service: service,
	}
}

type authHandler struct{
	service AuthService
}

type AuthHandler interface {
	SignUp(*gin.Context)
	SignIn(*gin.Context)
}

func (h *authHandler) SignUp(c *gin.Context){
	fmt.Println("Chegou no user Handler")

	var ur SignUpRequest

	if err := c.ShouldBindJSON(&ur); err != nil {
		errMessage := fmt.Sprintf("error when binding json: %s", err.Error())
		errJson := apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
		return
	}

	token, err := h.service.SignUp(ur);
	
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{Token: token})
}

func (h *authHandler) SignIn (c *gin.Context){
	var ur SignInRequest
	
	if err := c.ShouldBindJSON(&ur); err != nil{
		errMessage := fmt.Sprintf("error when binding json: %s", err.Error())
		errJson := apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
		return
	}

	token, err := h.service.SignIn(ur)

	if err != nil{
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, AuthResponse{Token: token})
}