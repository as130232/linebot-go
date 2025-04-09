package cmd

import "linebot-go/application/service/line"

func InitApp() *App {
	lineBotService := line.NewBotService()
	return NewApp(lineBotService)
}
