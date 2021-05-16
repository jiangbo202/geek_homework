/**
 * @Author: jiangbo
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2021/05/16 9:38 下午
 */

package routers

import (
	"github.com/gin-gonic/gin"
	v1 "jiang.geek/work04/api/app/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("user")
	{
		BaseRouter.POST("register", v1.Register)
		BaseRouter.POST("login", v1.Login)
	}
	return BaseRouter
}

