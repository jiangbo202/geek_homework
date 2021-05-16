/**
 * @Author: jiangbo
 * @Description:
 * @File:  AuthMiddleware
 * @Version: 1.0.0
 * @Date: 2021/05/16 9:33 下午
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"jiang.geek/work04/model/db_model"
	"jiang.geek/work04/utils/data"
	"net/http"
	"strings"
)

func AuthMiddlewar() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer "){
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"message": "权限不足",
			})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := data.ParseToken(tokenString)
		if err != nil || !token.Valid{
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"message": "权限不足",
			})
			context.Abort()
			return

		}
		// 验证通过后获取claims里面的userId
		userId := claims.UserId
		DB := data.GetDB()
		var user db_model.User
		DB.First(&user, userId)
		// 用户
		if user.ID == 0{
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"message": "用户不存在",
			})
			context.Abort()
			return
		}
		// 用户存在，将用户写入上下文
		context.Set("user", user)
		context.Next()
	}
}
