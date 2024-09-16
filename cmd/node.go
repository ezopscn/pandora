package cmd

import (
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"pandora/common"
	"pandora/initialize"
	"pandora/pkg/gedis"
	"pandora/pkg/utils"
	"strconv"
	"strings"
	"time"
)

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringVarP(&common.SystemConfigFilename, "config", "", common.SystemConfigFilename, "可选，指定服务启动配置文件")
}

// 节点信息命令
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "查看当前程序的相关节点信息",
}

// 节点状态
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "查看当前程序的相关节点状态信息",
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化
		initialize.Config()
		initialize.SystemLogger()
		initialize.AccessLogger()
		initialize.Redis()

		// 处理节点信息
		t := table.NewWriter()
		t.AppendHeader(table.Row{"序号", "节点UUID", "角色", "在线时长"})
		var rows []table.Row

		// 当前时间
		now := time.Now().Unix()

		// 获取节点信息
		rctx := context.Background()
		master, _ := common.RedisCache.Get(rctx, common.RK_MASTER_ID).Result()

		// 批量获取 Key
		keys, err := gedis.GetKeysWithPrefix(common.RedisCache, common.RKP_NODE_ID)

		if err == nil {
			// 逐个获取 key 的值
			for id, key := range keys {
				v, err := common.RedisCache.Get(rctx, key).Result()
				if err == nil {
					ts, _ := strconv.ParseInt(v, 10, 64)
					deviceId := strings.TrimPrefix(key, common.RKP_NODE_ID)
					if ts != 0 {
						var role = "worker"
						if deviceId == master {
							role = "master"
						}
						rows = append(rows, table.Row{id + 1, deviceId, role, utils.ConvertSecondsToHumanReadable(now - ts)})
					}
				} else {
					common.SystemLog.Error(fmt.Sprintf("获取 Redis Key %s 的值失败", key))
				}
			}
		}
		t.AppendRows(rows)
		fmt.Println(t.Render())
	},
}
