package prod

import (
	"linebot-go/manifest/config"
)

func CreateServerConfig() config.ServerConfig {
	return config.ServerConfig{
		HttpServer: &config.HttpServerConfig{Address: ":8081", Mode: "release"},
	}
}
