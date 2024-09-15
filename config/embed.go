package config

import (
	"embed"
)

//go:embed pandora.yaml
//go:embed version
var Fs embed.FS // 解决开发过程中指定配置启动问题
