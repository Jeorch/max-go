package maxconfig

import "github.com/alfredyang1986/blackmirror/bmconfighandle"

type PhForwardConfig struct {
	HostA string
	PortA string
	HostB string
	PortB string
}

func (mc *PhForwardConfig) GenerateConfig() {
	//TODO: 配置文件路径 待 用脚本指定dev路径和deploy路径
	configPath := "resource/forwardconfig.json"
	profileItems := bmconfig.BMGetConfigMap(configPath)

	mc.HostA = profileItems["HostA"].(string)
	mc.PortA = profileItems["PortA"].(string)
	mc.HostB = profileItems["HostB"].(string)
	mc.PortB = profileItems["PortB"].(string)

}
