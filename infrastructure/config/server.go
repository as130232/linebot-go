package config

import (
	"linebot-go/manifest/config"
	"linebot-go/manifest/config/local"
)

func NewServerConfig() *config.ServerConfig {
	var serverConfig config.ServerConfig
	//根據環境取得對應環境變數
	//appEnv := os.Getenv("APP_ENV")
	//switch appEnv {
	//case "local":
	//	serverConfig = local.CreateServerConfig()
	//case "dev":
	//	serverConfig = dev.CreateServerConfig()
	//case "prod":
	//	serverConfig = prod.CreateServerConfig()
	//default:
	//	panic("APP_ENV must be local|dev|cqa|uat|prod")
	//}
	serverConfig = local.CreateServerConfig()
	return &serverConfig
}
