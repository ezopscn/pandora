package initialize

import (
	"github.com/gin-gonic/gin"
	"pandora/common"
	"pandora/middleware"
	"pandora/route"
)

// 路由初始化
func Router() *gin.Engine {
	// 创建一个没中间件的路由引擎
	r := gin.New()

	// 中间件
	r.Use(middleware.AccessLog) // 请求日志中间件
	r.Use(middleware.Cors)      // 跨域访问中间件
	r.Use(middleware.Exception) // 异常捕获中间件

	rg := r.Group(common.SYSTEM_API_PREFIX)
	route.PublicRoutes(rg) // 开放路由组
	return r
}
