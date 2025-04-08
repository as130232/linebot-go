package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloHandler struct {
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (hh *HelloHandler) HelloWorld(c *gin.Context) {
	name := c.Query("name")
	result := "Hello, " + name + "!"
	c.JSON(http.StatusOK, result)
	c.Abort()
}
