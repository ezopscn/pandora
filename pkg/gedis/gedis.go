package gedis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// 批量获取以指定前缀开头的键
func GetKeysWithPrefix(rdb *redis.Client, prefix string) ([]string, error) {
	var keys []string
	var cursor uint64
	var rctx = context.Background()

	for {
		// SCAN 命令，游标为 0 从头开始
		newKeys, newCursor, err := rdb.Scan(rctx, cursor, prefix+"*", 0).Result()
		if err != nil {
			return nil, err
		}

		keys = append(keys, newKeys...)
		cursor = newCursor

		// 如果游标为 0 则表示遍历结束
		if cursor == 0 {
			break
		}
	}
	return keys, nil
}
