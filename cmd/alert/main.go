package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"pandora/common"
	"pandora/config"
	"pandora/initialize"
	"pandora/pkg/ticker"
	"pandora/pkg/utils"
	"strings"
	"time"
)

// 服务名称
var systemName = "pandora-alert"

func init() {
	// 指定配置文件参数
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&common.ConfigFileName, "config", "", common.ConfigFileName, "可选，指定服务启动配置文件")
	startCmd.Flags().StringVarP(&common.ListenAddress, "listen", "", common.ListenAddress, "可选，指定服务启动的监听地址")
	startCmd.Flags().StringVarP(&common.ListenPort, "port", "", "", "可选，指定服务启动的监听端口")

	// 查看版本信息
	rootCmd.AddCommand(infoCmd)
}

// 命令入口
var rootCmd = &cobra.Command{
	Use:   systemName,
	Short: fmt.Sprintf("Pandora（潘多拉），基于 Go + Gin + Gorm + Casbin 开发的监控系统后端（%s 端）", systemName),
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) { },
}

// 系统信息命令
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "查看当前程序的相关系统信息",
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		t.AppendHeader(table.Row{"名称", "信息"})
		t.AppendRows([]table.Row{
			{"项目名称", common.SystemName},
			{"系统版本", common.SystemVersion},
			{"GO 版本", common.SYSTEM_GO_VERSION},
			{"开发者", common.SYSTEM_DEVELOPER_NAME},
			{"开发者邮箱", common.SYSTEM_DEVELOPER_EMAIL},
		})
		fmt.Println(t.Render())
	},
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "参数化启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		// 判断用户是否传递监听地址
		var listenAddress string
		if common.ListenAddress != "" {
			if common.ListenAddress == "0.0.0.0" || utils.IsIPv4(common.ListenAddress) {
				listenAddress = common.ListenAddress
			} else {
				fmt.Println("传递的监听地址格式不合法，只支持 IPV4 或者 0.0.0.0")
				os.Exit(0)
			}
		}

		// 判断用户是否传递监听端口
		var listenPort string
		if common.ListenPort != "" {
			if utils.IsPort(common.ListenPort) {
				listenPort = common.ListenPort
			} else {
				fmt.Println("传递的监听端口格式不合法，只支持 1-65535 之间的数字")
				os.Exit(0)
			}
		}

		// 配置文件初始化
		initialize.Config()

		// 日志初始化
		common.SystemLog = initialize.NewLogger(common.Config.Alert.Log.System) // 系统日志初始化
		common.AccessLog = initialize.NewLogger(common.Config.Alert.Log.Access) // 访问日志初始化

		// 没设置参数则使用配置文件中的
		if listenAddress == "" {
			listenAddress = common.Config.Alert.System.Listen
		}

		// 没设置参数则使用配置文件中的
		if listenPort == "" {
			listenPort = common.Config.Alert.System.Port
		}

		// Logo
		fmt.Println(common.ALERT_LOGO)

		// 客户端标识
		id, _ := uuid.NewUUID()
		common.SystemUUID = id.String()
		common.SystemLog.Info("系统启动标识UUID：", common.SystemUUID)

		// 初始化 Redis 连接
		initialize.Redis()

		// 客户端启动时注册
		go ticker.HeartbeatTicker(common.RedisCache, common.SystemUUID)

		// 竞选 Master
		go ticker.MasterElectionTicker(common.RedisCache, common.SystemUUID)

		// 数据接口部分
		r := initialize.AlertRouter() // 路由初始化
		addr := fmt.Sprintf("%s:%s", listenAddress, listenPort)
		common.SystemLog.Info("服务的监听地址为：", addr)

		server := http.Server{
			Addr:    addr,
			Handler: r,
		}

		// 启动服务
		go func() {
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				common.SystemLog.Error(err.Error())
				panic(err)
			}
		}()

		// 接收优雅关闭信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		// 等待5秒然后停止服务
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			common.SystemLog.Error(err.Error())
			panic(err)
		}
		common.SystemLog.Info("服务正常的停止完成")
	},
}

// 所有子命令添加到 root 命令，输入 cmd 的入口
func execute() {
	// 初始化变量
	common.FS = config.Fs
	common.SystemName = systemName
	common.SystemTitle = strings.ToUpper(systemName)

	// 读取版本号
	version, err := common.FS.ReadFile(common.VersionFileName)
	if err != nil {
		panic(err)
	}

	// 设置全局版本号
	if string(version) != "" {
		common.SystemVersion = string(version)
	}

	if err = rootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}

// 主函数
func main() {
	execute()
}
