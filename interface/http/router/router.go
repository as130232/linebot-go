package router

import (
	"github.com/gin-gonic/gin"
	"linebot-go/cmd"
	"linebot-go/global"
	"linebot-go/interface/http/handler"
)

func InitRouter(app *cmd.App) *gin.Engine {
	router := gin.New()
	//swagger
	//docs.Swagger.Title = ""

	serviceRouter := router.Group("/" + global.AppName)
	lineHandler := handler.NewLineHandler(app.LineBotService)
	serviceRouter.POST("callback", lineHandler.Callback)

	helloHandler := handler.NewHelloHandler()
	serviceRouter.GET("/hello", helloHandler.HelloWorld)

	return router
}
