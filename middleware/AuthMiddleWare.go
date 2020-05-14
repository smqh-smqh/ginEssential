package middleware

import (
	"net/http"
	"strings"

	"example.com/ginEssential/common"
	"example.com/ginEssential/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenstring := c.GetHeader("Authorization")

		if tokenstring == "" || !strings.HasPrefix(tokenstring, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请求头错误"})
			c.Abort()
			return
		}

		tokenstring = tokenstring[7:]
		token, claims, err := common.ParseToken(tokenstring)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 402, "msg": "token错误"})
			c.Abort()
			return
		}

		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "msg": "用户不存在"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
