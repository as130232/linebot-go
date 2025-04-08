package line

import (
	_ "github.com/line/line-bot-sdk-go/v8/linebot"
	_ "github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	_ "github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type BotService struct {
}

func NewBotService() *BotService {
	return &BotService{}
}
