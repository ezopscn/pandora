package route

import (
	"github.com/gin-gonic/gin"
	"pandora/api/v1/public"
)

// 开放路由组
func ServerPublicRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/health", public.HealthHandler) // 健康检查接口
	rg.GET("/info", public.InfoHandler)     // 信息接口
	return rg
}
