package handler

import (
	"aeperez24/basego/port"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MiddlewareConfigGin struct {
	AuthenticationMiddleware func(*gin.Context)
}

type ServerConfigurationGin struct {
	MiddlewareConfigGin MiddlewareConfigGin
	Port                string
	HandlerConfig       HandlerConfigGin
}

type GinServerImpl struct {
	ServerConfigurationGin
	HttpServer http.Server
}

func NewGinServer(config ServerConfigurationGin) port.Server {
	return GinServerImpl{ServerConfigurationGin: config}
}

func (mserver GinServerImpl) Start() error {
	router := gin.Default()
	//authMiddleware := mserver.MiddlewareConfigGin.AuthenticationMiddleware
	router.POST("/user/signin", mserver.HandlerConfig.AuthenticationHandler.Authenticate)
	router.POST("/user/signup", mserver.HandlerConfig.UserHandler.CreateUser)

	mserver.HttpServer = http.Server{Addr: mserver.Port, Handler: router}

	return mserver.HttpServer.ListenAndServe()
}

func (mserver GinServerImpl) Stop() {
	mserver.HttpServer.Shutdown(context.Background())
}
