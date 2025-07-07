package middleware

import (
	"net/http"
	"strings"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/FamousLuisin/agoraspace/internal/handler/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenMiddleware(c *gin.Context){
	tokenRequest := c.Request.Header.Get("Authorization")

	if tokenRequest == "" || !strings.Contains(tokenRequest, "Bearer ") {
		errMessage := apperr.NewAppError("invalid token", apperr.ErrUnauthorized, http.StatusUnauthorized)
		c.JSON(errMessage.Code, errMessage)
		c.Abort()
		return
	}

	tokenValue := strings.TrimPrefix(tokenRequest, "Bearer ")

	token, err := auth.VerifyToken(tokenValue)

	if err != nil {
		errMessage := apperr.NewAppError(err.Error(), apperr.ErrUnauthorized, http.StatusUnauthorized)
		c.JSON(errMessage.Code, errMessage)
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Set("subject", claims["sub"])
}