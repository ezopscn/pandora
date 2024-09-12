package config

import (
	"embed"
)

//go:embed pandora.yaml
var Fs embed.FS // 解决开发过程中指定配置启动问题
