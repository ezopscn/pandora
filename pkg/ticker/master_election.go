package ticker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// Master 竞选
func TryToBecomeMaster(rdb *redis.Client, id string) {
	ctx := context.Background()
	key := "MASTER-ID"
	for {
		// 获取指定 Key 的 Value
		result, err := rdb.Get(ctx, key).Result()
		if err != nil {
			rdb.SetNX(ctx, key, id, time.Second*30) // 过期 30 秒
		}
		// 是当前节点，则更新
		if result == id {
			rdb.SetEx(ctx, key, id, time.Second*30) // 更新过期时间
		}
		time.Sleep(time.Second * 5) // 至少 5 次上报不成功，则 Master 被剔除
	}
}
