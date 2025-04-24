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
	lineBotChannelSecret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	lineBotChannelToken := os.Getenv("LINE_BOT_CHANNEL_TOKEN")
	dbUserName := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	return config.ServerConfig{
		HttpServer: &config.HttpServerConfig{Address: ":" + port, Mode: gin.ReleaseMode},
		DbConfig: &config.DbConfig{
			DbHost:   "aws-0-ap-southeast-1.pooler.supabase.com",
			DbPort:   6543,
			DbName:   "postgres",
			Username: dbUserName,
			Password: dbPassword,
		},
		LineConfig: &config.LineConfig{
			ChannelId:     "2007224382", // 宏甘
			ChannelSecret: lineBotChannelSecret,
			ChannelToken:  lineBotChannelToken,
		},
	}
}
