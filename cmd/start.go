package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"pandora/common"
	"pandora/initialize"
	"pandora/pkg/utils"
	"time"
)

func init() {
	// 指定配置文件参数
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&common.SystemConfigFilename, "config", "", common.SystemConfigFilename, "可选，指定服务启动配置文件")
	startCmd.Flags().StringVarP(&common.SystemListenAddress, "listen", "", common.SystemListenAddress, "可选，指定服务启动的监听地址")
	startCmd.Flags().StringVarP(&common.SystemListenPort, "port", "", common.SystemListenPort, "可选，指定服务启动的监听端口")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "参数化启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		// 配置文件初始化
		initialize.Config()

		// 判断用户是否传递监听地址和端口是否合法，没有传参，则使用配置文件中的配置
		if common.SystemListenAddress == "" {
			common.SystemListenAddress = common.Config.System.Listen
		}

		if common.SystemListenAddress != "0.0.0.0" || !utils.IsIPv4(common.SystemListenAddress) {
			fmt.Println("传递的监听地址格式不合法，只支持 IPV4 或者 0.0.0.0")
			os.Exit(0)
		}

		if common.SystemListenPort == "" {
			common.SystemListenPort = common.Config.System.Port
		}

		if !utils.IsPort(common.SystemListenPort) {
			fmt.Println("传递的监听端口格式不合法，只支持 1-65535 之间的数字")
			os.Exit(0)
		}

		// 初始化日志
		initialize.SystemLogger() // 系统日志初始化
		initialize.AccessLogger() // 访问日志初始化

		// Logo
		fmt.Println(common.LOGO)

		// 客户端标识
		id, _ := uuid.NewUUID()
		common.SystemUUID = id.String()
		common.SystemLog.Info("系统启动标识UUID：", common.SystemUUID)

		// 初始化 Redis 连接
		initialize.Redis()

		// 数据接口部分
		r := initialize.Router() // 路由初始化
		addr := fmt.Sprintf("%s:%s", common.SystemListenAddress, common.SystemListenPort)
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
