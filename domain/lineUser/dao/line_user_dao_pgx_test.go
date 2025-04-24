package dao

import (
	"context"
	"github.com/jackc/pgx/v5"
	"linebot-go/global"
	"linebot-go/infrastructure/pkg/database"
	"linebot-go/manifest/config/local"
	"log"
	"testing"
)

func getRepoPgx() *LineUserDaoPgx {
	serverConfig := local.CreateServerConfig()
	global.ServerConfig = &serverConfig
	database.SetupDBByPgx()
	//database.SetupDbByPgx()
	return NewLineUserDaoPgx()
}
func TestGetAllByPgx(t *testing.T) {
	repo := getRepoPgx()
	result, err := repo.FindAll()
	if err != nil {
		log.Fatalf("err: %+v", err)
	}
	log.Printf("result: %+v", result)

	defer func(SqlDB *pgx.Conn, ctx context.Context) {
		err := SqlDB.Close(ctx)
		if err != nil {

		}
	}(global.DbPgx, context.Background())
}
