package ticker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"pandora/common"
	"time"
)

// 周期性心跳上报
func HeartbeatTicker(rdb *redis.Client, id string) {
	rctx := context.Background()
	key := common.RKP_NODE_ID + id
	value := time.Now().Unix()
	expire := 10 * time.Second            // 过期时间，10s
	tk := time.NewTicker(5 * time.Second) // 上报间隔，5s
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			rdb.SetEx(rctx, key, value, expire)
		}
	}
}
