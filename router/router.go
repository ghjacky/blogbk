package router

import (
	"blogbk/middleware"
	"github.com/gin-gonic/gin"
	"strings"
)

var Router = gin.Default()

var apiPrefix = "/api/v1"

func InitRouter() {
	useMiddleware()
	categoryRouter()
	postRouter()
	tagRouter()
	userRouter()
}

func useMiddleware() {
	Router.Use(middleware.Auth())
}

func api(entry string) string {
	return strings.Join([]string{strings.TrimSuffix(apiPrefix, "/"), strings.Trim(entry, "/")}, "/")
}

type SHttpResp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newHttpResp(message string, data interface{}) interface{} {
	return SHttpResp{Message: message, Data: data}
}
