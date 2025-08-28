package forum

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/gin-gonic/gin"
)

func NewForumHandler(service ForumService) ForumHandler {
	return &forumHandler{
		service: service,
	}
}

type forumHandler struct {
	service ForumService
}

type ForumHandler interface {
	CreateForum(*gin.Context)
	GetAllForums(*gin.Context)
	GetForumById(c *gin.Context)
	UpdateForum(c *gin.Context)
	DeleteForum(c *gin.Context)
}

func (h *forumHandler) CreateForum(c *gin.Context){
	var f ForumRequest

	if err := c.ShouldBindJSON(&f); err != nil {
		errJson := apperr.NewAppError(fmt.Sprintf("error when binding json: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
		return
	}

	subject, ok := c.Get("subject")

	if !ok {
		errJson := apperr.NewAppError("error subject not found", apperr.ErrUnauthorized, http.StatusUnauthorized)
		c.JSON(errJson.Code, errJson)
		return
	}

	if err := h.service.CreateForum(f, subject.(string)); err != nil {
		c.JSON(err.Code, err)
		return
	}
}

func (h *forumHandler) GetAllForums(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "0")
	perPageStr := c.DefaultQuery("perPage", "10")

	page, _page := strconv.Atoi(pageStr)
	perPage, _perPage := strconv.Atoi(perPageStr)

	if err := errors.Join(_page, _perPage); err != nil {
		errJson := apperr.NewAppError(fmt.Sprintf("error converting query: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
	}
	
	forums, err := h.service.GetAllForums(page, perPage)

	if err != nil {
		c.JSON(err.Code, err)
	}

	c.JSON(http.StatusOK, forums)
}

func (h *forumHandler) GetForumById(c *gin.Context) {
	forumId := c.Param("id")
	
	forum, err := h.service.GetForumById(forumId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, forum)
}

func (h *forumHandler) UpdateForum(c *gin.Context) {
	var forumRequest ForumRequest

	if err := c.ShouldBindJSON(&forumRequest); err != nil {
		errJson := apperr.NewAppError(fmt.Sprintf("error when binding json: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
		c.JSON(errJson.Code, errJson)
		return
	}

	forumId := c.Param("id")
	
	identifier, _ := c.Get("subject")
	identifierStr := identifier.(string)

	forum, err := h.service.UpdateForum(forumRequest, identifierStr, forumId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, forum)
}

func (h *forumHandler) DeleteForum(c *gin.Context) {
	forumId := c.Param("id")
	
	identifier, _ := c.Get("subject")
	identifierStr := identifier.(string)

	if err := h.service.DeleteForum(identifierStr, forumId); err != nil {
		c.JSON(err.Code, err)
		return
	}
}