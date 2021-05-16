/**
 * @Author: jiangbo
 * @Description:
 * @File:  routers
 * @Version: 1.0.0
 * @Date: 2021/05/16 5:48 下午
 */

package configs

import (
	"github.com/gin-gonic/gin"
	"jiang.geek/work04/utils/middleware"
	"jiang.geek/work04/utils/routers"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	publicGroup := router.Group("api/v1")  // 注册基础功能路由 不做鉴权
	{
		routers.InitBaseRouter(publicGroup)
	}
	PrivateGroup := router.Group("api/v1")
	PrivateGroup.Use(middleware.AuthMiddlewar())  // 需要认证
	{
		routers.InitUserRouter(PrivateGroup)
	}
	return router
}

