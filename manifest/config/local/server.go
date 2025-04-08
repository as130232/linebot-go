package local

import (
	"github.com/gin-gonic/gin"
	"linebot-go/manifest/config"
	"os"
)

func CreateServerConfig() config.ServerConfig {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}
	return config.ServerConfig{
		HttpServer: &config.HttpServerConfig{Address: ":" + port, Mode: gin.ReleaseMode},
		LineConfig: &config.LineConfig{
			ChannelId:     "2007224382",
			ChannelSecret: "a0e83c3ca58f0ddc6c4157c7fa91ca28",
		},
	}
}
