package routes

import (
	"github.com/FamousLuisin/agoraspace/internal/handler/meta"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	{
		r.GET("/version", meta.Version)
	}
}