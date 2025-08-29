package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/FamousLuisin/agoraspace/internal/models"
	"github.com/FamousLuisin/agoraspace/internal/services"
	"github.com/gin-gonic/gin"
)

func NewAuthHandler(service services.AuthService) AuthHandler {
	return &authHandler{
		service: service,
	}
}

type authHandler struct{
	service services.AuthService
}

type AuthHandler interface {
	SignUp(*gin.Context)
	SignIn(*gin.Context)
}

func (h *authHandler) SignUp(c *gin.Context){
	var ur models.SignUpRequest

	if err := c.ShouldBindJSON(&ur); err != nil {
		errMessage := fmt.Sprintf("error when binding json: %s", err.Error())
		errJson := apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
		return
	}

	parsedBirthday, parse := time.Parse("2006-01-02", ur.BirthdayString)
	
	if parse != nil{
		errMessage := fmt.Sprintf("error when parse birthday: %s", parse.Error())
		errJson := apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
		return
	}

	ur.Birthday = parsedBirthday

	token, err := h.service.SignUp(ur);
	
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.SetCookie("agoraToken", token, 900, "/", "localhost", false, true)
	c.JSON(http.StatusCreated, models.AuthResponse{Token: token})
}

func (h *authHandler) SignIn(c *gin.Context){
	var ur models.SignInRequest
	
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

	c.SetCookie("agoraToken", token, 900, "/", "localhost", false, true)
	c.JSON(http.StatusOK, models.AuthResponse{Token: token})
}