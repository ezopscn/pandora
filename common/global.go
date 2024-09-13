package common

import (
	"embed"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局工具
var (
	FS         embed.FS           // 文件打包
	SystemLog  *zap.SugaredLogger // 系统日志工具
	AccessLog  *zap.SugaredLogger // 访问日志工具
	MySQLDB    *gorm.DB           // 数据库连接
	RedisCache *redis.Client      // 缓存连接
)

// 全局常量
const (
	MSEC_TIME_FORMAT       = "2006-01-02 15:04:05.000" // 时间格式化格式
	SEC_TIME_FORMAT        = "2006-01-02 15:04:05"     // 时间格式化格式
	DATE_TIME_FORMAT       = "2006-01-02"              // 时间格式化格式
	SYSTEM_DEVELOPER_NAME  = "DK"                      // 开发者
	SYSTEM_DEVELOPER_EMAIL = "ezops.cn@gmail.com"      // 邮箱
	SYSTEM_GO_VERSION      = "1.23.0"                  // Go 版本
)

// 系统可变配置
var (
	ConfigFileName  = "pandora.yaml" // 配置文件
	VersionFileName = "version"      // 版本文件
	ListenAddress   = "0.0.0.0"      // 监听地址
	ListenPort      = ""             // 监听端口
	SystemName      = ""             // 系统名称
	SystemVersion   = ""             // 系统版本
)
