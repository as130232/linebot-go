package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"linebot-go/global"
	"log"
	"time"
)

func SetupDBByPgx() {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		global.ServerConfig.DbConfig.Username, global.ServerConfig.DbConfig.Password,
		global.ServerConfig.DbConfig.DbHost, global.ServerConfig.DbConfig.DbPort, global.ServerConfig.DbConfig.DbName)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	log.Println("Connect db success by pgx. version:", version)
	global.DbPgx = conn
}

func SetupDBByGorm() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		global.ServerConfig.DbConfig.DbHost, global.ServerConfig.DbConfig.Username, global.ServerConfig.DbConfig.Password,
		global.ServerConfig.DbConfig.DbName, global.ServerConfig.DbConfig.DbPort)
	// 確保只初始化一次
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(global.ServerConfig.DbConfig.LogMode),
	})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetMaxOpenConns(150) //设置数据库连接池最大连接数
	sqlDB.SetConnMaxLifetime(60 * time.Second)

	log.Println("Connect db success by gorm.")
	global.DbGorm = db
}
