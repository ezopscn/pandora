package cmd

import (
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"pandora/common"
	"pandora/initialize"
	"strings"
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
		initialize.Config()
		initialize.SystemLogger()
		initialize.AccessLogger()
		initialize.Redis()

		// 获取节点信息
		var workers []string
		rctx := context.Background()
		master, _ := common.RedisCache.Get(rctx, common.RK_MASTER_ID).Result()
		keys, _, _ := common.RedisCache.Scan(rctx, 0, common.RKP_NODE_ID+"*", 0).Result()
		for _, key := range keys {
			workers = append(workers, strings.TrimPrefix(key, common.RKP_NODE_ID))
		}

		// 处理节点信息
		t := table.NewWriter()
		t.AppendHeader(table.Row{"序号", "节点UUID", "角色"})
		var rows []table.Row
		for id, worker := range workers {
			id += 1
			var role = "worker"
			if worker == master {
				role = "master"
			}
			rows = append(rows, table.Row{id, worker, role})
		}
		t.AppendRows(rows)
		fmt.Println(t.Render())
	},
}
