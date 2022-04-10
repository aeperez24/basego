package handler

import (
	"aeperez24/basego/config"
	"aeperez24/basego/middleware"
	"aeperez24/basego/model"
	"aeperez24/basego/services"
	"aeperez24/basego/usercase"
)

func BuildServerConfigGin(portNumber string, tokenKey string, mongo config.MongoCofig) ServerConfigurationGin {
	tokenService := services.NewTokenService(tokenKey)
	userRepo := model.NewUserMongoRepository(mongo)
	userService := services.NewUserService(userRepo)
	authHandler := NewAuthenticationHandlerGin(userService, tokenService)

	userUserCase := usercase.UserUsercase{
		UserService: userService}

	userHandler := NewGinUserhandler(userUserCase)

	handlerConfig := HandlerConfigGin{
		AuthenticationHandler: authHandler,
		UserHandler:           userHandler,
	}

	middlewareConfig := MiddlewareConfigGin{AuthenticationMiddleware: middleware.NewAuthenticationMiddlwareGin(tokenService)}
	serverConfig := ServerConfigurationGin{
		Port:                ":" + portNumber,
		HandlerConfig:       handlerConfig,
		MiddlewareConfigGin: middlewareConfig,
	}

	return serverConfig
}
