package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/common"
	"pandora/pkg/response"
)

// 健康检测接口
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// 系统信息接口
func InfoHandler(ctx *gin.Context) {
	response.SuccessWithData(gin.H{
		"name":         common.SYSTEM_NAME,
		"chinese_name": common.SYSTEM_CHINESE_NAME,
		"desc":         common.SYSTEM_DESC,
		"version":      common.SystemVersion,
		"go":           common.SYSTEM_GO_VERSION,
		"developer":    common.SYSTEM_DEVELOPER_NAME,
		"email":        common.SYSTEM_DEVELOPER_EMAIL,
	})
}
