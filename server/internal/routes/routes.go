package routes

import (
	"github.com/FamousLuisin/agoraspace/internal/db"
	appAuth "github.com/FamousLuisin/agoraspace/internal/handler/auth"
	"github.com/FamousLuisin/agoraspace/internal/handler/meta"
	"github.com/FamousLuisin/agoraspace/internal/handler/user"
	"github.com/FamousLuisin/agoraspace/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, db *db.Database){
	authPath := r.Group("/auth")
	protectedPath := r.Group("/api", middleware.VerifyTokenMiddleware)
	
	r.GET("/version", meta.Version)

	userRepository := user.NewUserRepository(db)
	authService := appAuth.NewAuthService(userRepository)
	authHandler := appAuth.NewAuthHandler(authService)
	
	authPath.POST("/signup", authHandler.SignUp)
	authPath.POST("/signin", authHandler.SignIn)

	protectedPath.GET("/version", meta.Version)
}