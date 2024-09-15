package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"pandora/common"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

// 系统信息命令
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "查看当前程序的相关系统信息",
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		t.AppendHeader(table.Row{"名称", "信息"})
		t.AppendRows([]table.Row{
			{"英文名称", common.SYSTEM_NAME},
			{"中文名称", common.SYSTEM_CHINESE_NAME},
			{"项目说明", common.SYSTEM_DESC},
			{"系统版本", common.SystemVersion},
			{"GO 版本", common.SYSTEM_GO_VERSION},
			{"开发人员", common.SYSTEM_DEVELOPER_NAME},
			{"开发邮箱", common.SYSTEM_DEVELOPER_EMAIL},
		})
		fmt.Println(t.Render())
	},
}
