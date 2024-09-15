package ticker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"pandora/common"
	"time"
)

// 周期性心跳上报
func HeartbeatTicker(rdb *redis.Client, id string) {
	ctx := context.Background()
	key := common.SystemTitle + ":UUID:" + id
	value := "online"
	expire := 10 * time.Second            // 过期时间，10s
	tk := time.NewTicker(5 * time.Second) // 上报间隔，5s
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			rdb.SetEx(ctx, key, value, expire)
		}
	}
}
