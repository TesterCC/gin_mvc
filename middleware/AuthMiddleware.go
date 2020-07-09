package middleware

import (
	"ginessential/common"
	"ginessential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// use for protect user info interface
func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// get authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer"){    // Bearer请求头，是OAuth2.0的规范，由前端请求时附带的Bearer
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":401, "msg":"权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]  // 避开前缀 Bearer:
		// 解析失败或者解析的token无效也返回权限不足
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":401, "msg":"权限不足"})
			ctx.Abort()
			return
		}

		// 通过验证后获取claims中的UserId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userId)

		// Verify User
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code":401, "msg":"权限不足"})
			ctx.Abort()
			return
		}

		// If user exist, write user info into context
		ctx.Set("user", user)
		ctx.Next()  // Next should be used only inside middleware.
	}
}