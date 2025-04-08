package local

import "linebot-go/manifest/config"

func CreateServerConfig() config.ServerConfig {
	return config.ServerConfig{
		HttpServer: &config.HttpServerConfig{Address: ":8081", Mode: "release"},
		LineConfig: &config.LineConfig{
			ChannelId:     "2007224382",
			ChannelSecret: "a0e83c3ca58f0ddc6c4157c7fa91ca28",
		},
	}
}
