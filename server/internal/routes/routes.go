package routes

import (
	"github.com/FamousLuisin/agoraspace/internal/db"
	appAuth "github.com/FamousLuisin/agoraspace/internal/handler/auth"
	"github.com/FamousLuisin/agoraspace/internal/handler/forum"
	"github.com/FamousLuisin/agoraspace/internal/handler/meta"
	"github.com/FamousLuisin/agoraspace/internal/handler/user"
	"github.com/FamousLuisin/agoraspace/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, db *db.Database){
	authPath := r.Group("/auth")
	protectedPath := r.Group("/api", middleware.VerifyCookieTokenMiddleware, middleware.VerifyTokenMiddleware)
	
	r.GET("/version", meta.Version)

	userRepository := user.NewUserRepository(db)
	authService := appAuth.NewAuthService(userRepository)
	authHandler := appAuth.NewAuthHandler(authService)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)
	forumRepository := forum.NewForumRepository(db)
	forumService := forum.NewForumService(forumRepository)
	
	forumHandler := forum.NewForumHandler(forumService)
	
	authPath.POST("/signup", authHandler.SignUp)
	authPath.POST("/signin", authHandler.SignIn)

	protectedPath.GET("/user", userHandler.GetUsers)
	protectedPath.GET("/user/:username", userHandler.GetUserByUsername)
	protectedPath.PUT("/user/:username", userHandler.UpdateUser)
	protectedPath.DELETE("/user/:username", userHandler.DeleteUser)

	protectedPath.POST("/forum", forumHandler.CreateForum)
	protectedPath.GET("/forums", forumHandler.GetAllForums)
	protectedPath.GET("/forum/:id", forumHandler.GetForumById)
	protectedPath.PUT("/forum/:id", forumHandler.UpdateForum)
	protectedPath.DELETE("/forum/:id", forumHandler.DeleteForum)

	protectedPath.GET("/version", meta.Version)
}