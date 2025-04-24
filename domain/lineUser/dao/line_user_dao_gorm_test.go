package dao

import (
	"linebot-go/global"
	"linebot-go/infrastructure/pkg/database"
	"linebot-go/manifest/config/local"
	"log"
	"testing"
)

func getRepoGorm() *LineUserDaoGorm {
	serverConfig := local.CreateServerConfig()
	global.ServerConfig = &serverConfig
	database.SetupDBByGorm()
	return NewLineUserDaoGorm()
}

func TestFindAllByGorm(t *testing.T) {
	repo := getRepoGorm()
	result, err := repo.FindAll()
	if err != nil {
		log.Fatalf("err: %+v", err)
	}
	log.Printf("result: %+v", result)
}

func TestFindOneByGorm(t *testing.T) {
	repo := getRepoGorm()
	id := "ff52a57f7e6ba861c05be8837bfbcf0c6"
	result, err := repo.FindOne(id)
	if err != nil {
		log.Fatalf("err: %+v", err)
	}
	log.Printf("result: %+v", result)
}
