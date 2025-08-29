package handler

import (
	"github.com/gin-gonic/gin"
)

type version struct{
	Version string `json:"version"`
}

func Version(c *gin.Context){
	v := version{
		Version: "1.0.0",
	}
	c.JSON(200, v)
}