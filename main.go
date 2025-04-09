package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"linebot-go/cmd"
	"linebot-go/global"
	"linebot-go/infrastructure/config"
	"linebot-go/interface/http/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>

	// start
	global.AppName = "lineBot"
	global.ServerConfig = config.NewServerConfig()
	app := cmd.InitApp()
	ginRouter := router.InitRouter(app)
	InitHttpServer(ginRouter)

}

func InitHttpServer(ginRouter *gin.Engine) {
	httpServer := &http.Server{
		Addr:    global.ServerConfig.HttpServer.Address,
		Handler: ginRouter,
	}
	// 設置信號通道
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// 啟動 HTTP 伺服器
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("ListenAndServe: err: %v\n", err)
			panic(err)
		}
	}()
	// 等待接收信號
	<-sigs
	fmt.Println("Shutting down server...")

	// 創建帶有超時的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 優雅關閉伺服器
	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}
}
