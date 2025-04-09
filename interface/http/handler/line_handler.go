package handler

import (
	"github.com/gin-gonic/gin"
	"linebot-go/application/service/line"
)

type LineHandler struct {
	lineBotService *line.BotService
}

func NewLineHandler(lineBotService *line.BotService) *LineHandler {
	return &LineHandler{lineBotService: lineBotService}
}

func (ll *LineHandler) Callback(c *gin.Context) {
	ll.lineBotService.CallbackHandler(c)
}
