package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/look4suman/events-api/routes/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not token sent"})
		slog.Error("Token not sent")
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		slog.Error("Not Authorized", "error", err)
		return
	}

	ctx.Set("UserId", userId)
	ctx.Next()
}
