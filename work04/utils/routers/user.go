/**
 * @Author: jiangbo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2021/05/16 5:51 下午
 */

package routers

import (
	"github.com/gin-gonic/gin"
	v1 "jiang.geek/work04/api/app/v1"
)

func InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("info", v1.Info)
	}
	return userRouter


}
