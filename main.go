package main

import (
	"embed"
	"pandora/cmd"
	"pandora/common"
)

//go:embed config/*
var fs embed.FS // 固定格式，打包的时候会将 config 目录下面的文件都一起打包

func main() {
	// 设置全局使用
	common.FS = fs

	// 读取版本号
	version, err := common.FS.ReadFile(common.SYSTEM_VERSION_FILENAME)
	if err != nil {
		panic(err)
	}

	// 设置全局版本号
	if string(version) != "" {
		common.SystemVersion = string(version)
	}

	cmd.Execute()
}
