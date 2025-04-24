package config

import (
	"gorm.io/gorm/logger"
	"time"
)

type ServerConfig struct {
	AppEnv     string
	HttpServer *HttpServerConfig
	LineConfig *LineConfig
	DbConfig   *DbConfig
}

type LineConfig struct {
	ChannelId     string
	ChannelSecret string
	ChannelToken  string
}
type HttpServerConfig struct {
	Address    string
	ServerName string
	Mode       string
}

type DbConfig struct {
	Username        string
	Password        string `json:"-"`
	DbHost          string
	DbPort          int
	DbName          string
	ConnMaxLifetime time.Duration
	LogMode         logger.LogLevel
}
