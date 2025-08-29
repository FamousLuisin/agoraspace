package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/FamousLuisin/agoraspace/internal/models"
	"github.com/FamousLuisin/agoraspace/internal/services"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(service services.UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}

type userHandler struct {
	service services.UserService
}

type UserHandler interface {
	GetMe(c *gin.Context)
	GetUserByUsername(*gin.Context)
	GetUsers(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

func (h *userHandler) GetUserByUsername(c *gin.Context){
	username := c.Param("username")
	user, err := h.service.GetUserByUsername(username)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUsers(c *gin.Context){
	pageStr := c.DefaultQuery("page", "0")
	perPageStr := c.DefaultQuery("perPage", "10")
	
	page, _page := strconv.Atoi(pageStr)
	perPage, _perPage := strconv.Atoi(perPageStr)
	
	if err := errors.Join(_page, _perPage); err != nil {
		errMessage := fmt.Sprintf("error converting query: %s", err.Error())
		errJson := apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
	}

	ur, err := h.service.GetUsers(page, perPage)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, ur)
}

func (h *userHandler) UpdateUser(c *gin.Context){
	var u models.UserDTO

	if err := c.ShouldBindJSON(&u); err != nil {
		errMessage := fmt.Sprintf("error when binding json: %s", err.Error())
		errJson := apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
		return
	}

	identifier, _ := c.Get("subject")
	identifierStr := identifier.(string)
	username := c.Param("username")

	if err := h.service.UpdateUser(u, identifierStr, username); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *userHandler) DeleteUser(c *gin.Context){
	identifier, _ := c.Get("subject")
	identifierStr := identifier.(string)
	username := c.Param("username")

	if err := h.service.DeleteUser(identifierStr, username); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func (h *userHandler) GetMe(c *gin.Context){
	identifier, _ := c.Get("subject")
	identifierStr := identifier.(string)

	user, err := h.service.GetUserById(identifierStr)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, user)
}