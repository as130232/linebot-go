package dev

import (
	"linebot-go/manifest/config"
	"os"
)

func CreateServerConfig() config.ServerConfig {
	port := os.Getenv("PORT")
	return config.ServerConfig{
		HttpServer: &config.HttpServerConfig{Address: ":" + port, Mode: "release"},
	}
}
