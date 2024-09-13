package ticker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// 心跳探测
func HeartbeatTicker(rdb *redis.Client, keyPrefix string, id string) {
	tk := time.NewTicker(5 * time.Second) // 心跳间隔，5s
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			rdb.SetEx(context.Background(), keyPrefix+":"+id, "online", 10*time.Second) // 更新 TTL
		}
	}
}
