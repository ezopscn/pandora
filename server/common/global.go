package common

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Logo 图形生成网站：http://patorjk.com/software/taag/
var Logo = `
 ██▓███   ▄▄▄       ███▄    █ ▓█████▄  ▒█████   ██▀███   ▄▄▄      
▓██░  ██▒▒████▄     ██ ▀█   █ ▒██▀ ██▌▒██▒  ██▒▓██ ▒ ██▒▒████▄    
▓██░ ██▓▒▒██  ▀█▄  ▓██  ▀█ ██▒░██   █▌▒██░  ██▒▓██ ░▄█ ▒▒██  ▀█▄  
▒██▄█▓▒ ▒░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█▄   ▌▒██   ██░▒██▀▀█▄  ░██▄▄▄▄██ 
▒██▒ ░  ░ ▓█   ▓██▒▒██░   ▓██░░▒████▓ ░ ████▓▒░░██▓ ▒██▒ ▓█   ▓██▒
▒▓▒░ ░  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒  ▒▒▓  ▒ ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░ ▒▒   ▓▒█░
░▒ ░       ▒   ▒▒ ░░ ░░   ░ ▒░ ░ ▒  ▒   ░ ▒ ▒░   ░▒ ░ ▒░  ▒   ▒▒ ░
░░         ░   ▒      ░   ░ ░  ░ ░  ░ ░ ░ ░ ▒    ░░   ░   ░   ▒   
               ░  ░         ░    ░        ░ ░     ░           ░  ░
                               ░                                  
`

// 版本信息
var Version = "1.0"

// 运行相关参数配置
var RunConfig = "config/pandora.yaml" // 运行配置文件

// 时间格式化
var (
	MsecTimeFormat = "2006-01-02 15:04:05.000"
	SecTimeFormat  = "2006-01-02 15:04:05"
	DateTimeFormat = "2006-01-02"
)

// 全局工具
var (
	Log *zap.SugaredLogger // 日志工具
	DB  *gorm.DB           // 数据库连接
)
