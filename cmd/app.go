package cmd

import "linebot-go/application/service/line"

type App struct {
	LineBotService *line.BotService
}

func NewApp(lineBotService *line.BotService) *App {
	return &App{
		LineBotService: lineBotService,
	}
}
