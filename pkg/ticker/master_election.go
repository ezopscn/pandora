package ticker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"pandora/common"
	"time"
)

// Master 竞选
func MasterElectionTicker(rdb *redis.Client, id string) {
	ctx := context.Background()
	key := common.SystemTitle + ":MASTER-ID"
	expire := time.Second * 30 // 过期 30 秒，意味着 Master 角色切换需要 30s
	for {
		// 获取指定 Key 的 Value
		r1, err1 := rdb.Get(ctx, key).Result()
		if err1 != nil {
			r2, _ := rdb.SetNX(ctx, key, id, expire).Result()
			if r2 {
				common.IsAlertMaster = true // 设置一个标识，用于其它判断
			}
		} else {
			// 如果当前节点是 Master 标识，但是 id 却不匹配，则退出当前程序
			if common.IsAlertMaster {
				if r1 != id {
					common.SystemLog.Error("当前节点已经不是 Master，但还是占用 Master 角色，所以退出程序")
					os.Exit(1)
				} else {
					// 是 Master，且能够对上 ID，则更新过期时间
					rdb.SetEx(ctx, key, id, expire)
				}
			}
		}

		time.Sleep(time.Second * 5) // 至少 5 次上报不成功，则 Master 被剔除
	}
}
