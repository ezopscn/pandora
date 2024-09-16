package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora/common"
	"pandora/dto"
	"pandora/pkg/gedis"
	"pandora/pkg/response"
	"pandora/pkg/utils"
	"strconv"
	"strings"
	"time"
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
	var nodes []dto.NodeStatus
	now := time.Now().Unix()
	rctx := context.Background()
	master, _ := common.RedisCache.Get(rctx, common.RK_MASTER_ID).Result()

	// 批量获取 Key
	keys, err := gedis.GetKeysWithPrefix(common.RedisCache, common.RKP_NODE_ID)
	if err == nil {
		for _, key := range keys {
			v, err1 := common.RedisCache.Get(rctx, key).Result()
			if err1 == nil {
				ts, _ := strconv.ParseInt(v, 10, 64)
				deviceId := strings.TrimPrefix(key, common.RKP_NODE_ID)
				if ts != 0 {
					var role = "worker"
					if deviceId == master {
						role = "master"
					}
					nodes = append(nodes, dto.NodeStatus{
						NodeName: deviceId,
						NodeRole: role,
						LiveTime: utils.ConvertSecondsToHumanReadable(now - ts),
					})
				}
			} else {
				common.SystemLog.Error(fmt.Sprintf("获取 Redis Key %s 的值失败", key))
			}
		}
	}
	response.SuccessWithData(gin.H{
		"nodes": nodes,
	})
}
