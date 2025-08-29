package routes

import (
	"github.com/FamousLuisin/agoraspace/internal/db"
	"github.com/FamousLuisin/agoraspace/internal/handler"
	"github.com/FamousLuisin/agoraspace/internal/middleware"
	"github.com/FamousLuisin/agoraspace/internal/repository"
	"github.com/FamousLuisin/agoraspace/internal/services"
	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	userH handler.UserHandler
	forumH handler.ForumHandler
	authH handler.AuthHandler
	memberH handler.MemberHandler
}

func BuildHandler(db *db.Database) *AppHandler {
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authService := services.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	forumRepository := repository.NewForumRepository(db)
	forumService := services.NewForumService(forumRepository)
	forumHandler := handler.NewForumHandler(forumService)

	memberRepository := repository.NewMemberRepository(db)
	memberService := services.NewMemberService(memberRepository, forumRepository)
	memberHandler := handler.NewMamberHandler(memberService)

	return &AppHandler{
		userH: userHandler,
		forumH: forumHandler,
		authH: authHandler,
		memberH: memberHandler,
	}
}

func InitRoutes(r *gin.RouterGroup, db *db.Database){
	handlers := BuildHandler(db)
	
	authPath := r.Group("/auth")
	protectedPath := r.Group("/api", middleware.VerifyCookieTokenMiddleware, middleware.VerifyTokenMiddleware)
	
	r.GET("/version", handler.Version)
	
	authPath.POST("/signup", handlers.authH.SignUp)
	authPath.POST("/signin", handlers.authH.SignIn)

	protectedPath.GET("/user", handlers.userH.GetUsers)
	protectedPath.GET("/user/:username", handlers.userH.GetUserByUsername)
	protectedPath.GET("/user/me", handlers.userH.GetMe)
	protectedPath.PUT("/user/:username", handlers.userH.UpdateUser)
	protectedPath.DELETE("/user/:username", handlers.userH.DeleteUser)

	protectedPath.POST("/forum", handlers.forumH.CreateForum)
	protectedPath.GET("/forums", handlers.forumH.GetAllForums)
	protectedPath.GET("/forum/:id", handlers.forumH.GetForumById)
	protectedPath.PUT("/forum/:id", handlers.forumH.UpdateForum)
	protectedPath.DELETE("/forum/:id", handlers.forumH.DeleteForum)

	protectedPath.POST("/joinForum/:id", handlers.memberH.JoinForum)
	protectedPath.DELETE("/leaveForum/:id", handlers.memberH.LeaveForum)
	protectedPath.GET("/member/forums/me", handlers.memberH.MyForums)
	protectedPath.GET("/member/forums/:id", handlers.memberH.ForumsMemberId)


	protectedPath.GET("/version", handler.Version)
}