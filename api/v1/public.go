package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/common"
	"pandora/pkg/response"
	"strings"
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

// 节点信息接口
func NodeStatusHandler(ctx *gin.Context) {
	var workers []string
	rctx := context.Background()
	master, _ := common.RedisCache.Get(rctx, common.RK_MASTER_ID).Result()
	keys, _, _ := common.RedisCache.Scan(rctx, 0, common.RKP_NODE_ID+"*", 0).Result()
	for _, key := range keys {
		node := strings.TrimPrefix(key, common.RKP_NODE_ID)
		if node != master {
			workers = append(workers, node)
		}
	}
	response.SuccessWithData(gin.H{
		"master":  master,
		"workers": workers,
	})
}
