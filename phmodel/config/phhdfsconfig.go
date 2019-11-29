package maxconfig

import "github.com/alfredyang1986/blackmirror/bmconfighandle"

type PhHdfsConfig struct {
	Host string
	Port string
}

func (mc *PhHdfsConfig) GenerateConfig() {
	//TODO: 配置文件路径 待 用脚本指定dev路径和deploy路径
	configPath := "resource/hdfsconfig.json"
	profileItems := bmconfig.BMGetConfigMap(configPath)

	mc.Host = profileItems["Host"].(string)
	mc.Port = profileItems["Port"].(string)

}
