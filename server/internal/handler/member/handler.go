package member

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewMamberHandler(service MemberService) MemberHandler {
	return &memberHandler{
		service: service,
	}
}

type memberHandler struct {
	service MemberService
}

type MemberHandler interface {
	JoinForum(c *gin.Context)
	LeaveForum(c *gin.Context)
	ForumsMemberId(c *gin.Context)
	MyForums(c *gin.Context)
}

func (h *memberHandler) JoinForum(c *gin.Context) {
	identifier, _ := c.Get("subject")
	identifierStr := identifier.(string)

	forumId := c.Param("id")

	if err := h.service.JoinForum(identifierStr, forumId); err != nil {
		c.JSON(err.Code, err)
		return
	}
}

func (h *memberHandler) LeaveForum(c *gin.Context){
	identifier, _ := c.Get("subject")
	identifierStr := identifier.(string)

	forumId := c.Param("id")

	if err := h.service.LeaveForum(identifierStr, forumId); err != nil {
		c.JSON(err.Code, err)
		return
	}
}

func (h *memberHandler) MyForums(c *gin.Context){
	identifier, _ := c.Get("subject")
	identifierStr := identifier.(string)

	members, err := h.service.FindForumsByMember(identifierStr)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, members)
}

func (h *memberHandler) ForumsMemberId(c *gin.Context){
	identifier := c.Param("id")

	members, err := h.service.FindForumsByMember(identifier)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, members)
}