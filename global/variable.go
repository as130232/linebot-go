package global

import (
	"github.com/jackc/pgx/v5"
	"gorm.io/gorm"
	"linebot-go/manifest/config"
)

var (
	AppName      string
	ServerConfig *config.ServerConfig
	DbPgx        *pgx.Conn
	DbGorm       *gorm.DB
)
