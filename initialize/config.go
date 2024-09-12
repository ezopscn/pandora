package initialize

import (
	"bytes"
	"github.com/spf13/viper"
	"log"
	"os"
	"pandora/common"
	"pandora/pkg/utils"
)

// 初始化配置
func Config() {
	// 读取的数据
	var bs []byte
	var err error

	// Viper 读取文件
	v := viper.New()
	v.SetConfigType("yaml")

	// 优先读取本地文件，然后才是 embed 打包的配置
	filename := common.ConfigFileName
	exist := utils.FileExist(filename)
	if exist {
		bs, err = os.ReadFile(filename)
	} else {
		bs, err = common.FS.ReadFile(filename)
	}
	if err != nil {
		panic(err)
	}

	// 解析配置
	err = v.ReadConfig(bytes.NewReader(bs))
	if err != nil {
		panic(err)
	}

	// 配置放到内存中
	settings := v.AllSettings()
	for i, setting := range settings {
		v.Set(i, setting)
	}

	// 设置全局变量，方便调用
	err = v.Unmarshal(&common.Config)
	if err != nil {
		panic(err)
	}

	log.Println("配置文件初始化完成：", filename)
}
