package router

import (
	"github.com/gin-gonic/gin"
	"linebot-go/global"
	"linebot-go/interface/http/handler"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	//swagger
	//docs.Swagger.Title = ""

	serviceRouter := router.Group("/" + global.AppName)

	helloHandler := handler.NewHelloHandler()
	serviceRouter.GET("/hello", helloHandler.HelloWorld)

	return router
}
