package routes

import (
	"github.com/FamousLuisin/agoraspace/internal/db"
	appAuth "github.com/FamousLuisin/agoraspace/internal/handler/auth"
	"github.com/FamousLuisin/agoraspace/internal/handler/forum"
	"github.com/FamousLuisin/agoraspace/internal/handler/member"
	"github.com/FamousLuisin/agoraspace/internal/handler/meta"
	"github.com/FamousLuisin/agoraspace/internal/handler/user"
	"github.com/FamousLuisin/agoraspace/internal/middleware"
	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	userH user.UserHandler
	forumH forum.ForumHandler
	authH appAuth.AuthHandler
	memberH member.MemberHandler
}

func BuildHandler(db *db.Database) *AppHandler {
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	authService := appAuth.NewAuthService(userRepository)
	authHandler := appAuth.NewAuthHandler(authService)

	forumRepository := forum.NewForumRepository(db)
	forumService := forum.NewForumService(forumRepository)
	forumHandler := forum.NewForumHandler(forumService)

	memberRepository := member.NewMemberRepository(db)
	memberService := member.NewMemberService(memberRepository, forumRepository)
	memberHandler := member.NewMamberHandler(memberService)

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
	
	r.GET("/version", meta.Version)
	
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


	protectedPath.GET("/version", meta.Version)
}