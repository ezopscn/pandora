package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"pandora/common"
)

// 命令入口
var rootCmd = &cobra.Command{
	Use:   common.SYSTEM_NAME,
	Short: fmt.Sprintf("%s（%s），%s", common.SYSTEM_TITLE, common.SYSTEM_CHINESE_NAME, common.SYSTEM_DESC),
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) { },
}

// 所有子命令添加到 root 命令，输入 cmd 的入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
