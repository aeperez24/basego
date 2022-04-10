package handler

import (
	"aeperez24/basego/config"
	"aeperez24/basego/middleware"
	"aeperez24/basego/model"
	"aeperez24/basego/services"
	"aeperez24/basego/usercase"
)

func BuildServerConfigGin(portNumber string, tokenKey string, mongo config.MongoCofig) ServerConfigurationGin {
	repo := model.NewAccountMongoRepository(mongo)
	tokenService := services.NewTokenService(tokenKey)
	userRepo := model.NewUserMongoRepository(mongo)
	userService := services.NewUserService(userRepo)

	accountUseCase := usercase.AccountUsercase{
		AccountRepository:  repo,
		TransactionService: services.NewTransactionService(repo),
	}
	accountHandler := GinAccountHandlerImpl{accountUseCase}
	authHandler := NewAuthenticationHandlerGin(userService, tokenService)

	userUserCase := usercase.UserUsercase{AccountRepository: repo,
		UserService: userService}

	userHandler := NewGinUserhandler(userUserCase)

	handlerConfig := HandlerConfigGin{
		AccountHandler:        accountHandler,
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
