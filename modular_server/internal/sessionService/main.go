package main

import (
	"common/config"
	"log"
	"sessionService/handler"
	"sessionService/repository"
	"sessionService/service"
	"sessionService/transport"
)

func main() {
	// _, _, sessionRepo, _ := repository.NewMockRepositoryImpl()
	redisCfg, err := config.LoadRedisConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	sessionRepo, err := repository.NewSessionRepositoryImpl(redisCfg)
	if err != nil {
		log.Fatal(err)
	}
	// sessionRepo, err := repository.NewSessionRepositoryImpl()
	sessionService, err := service.NewSessionServiceImpl(sessionRepo)
	if err != nil {
		log.Fatal(err)
	}
	grpcHandler, err := handler.NewSessionGRPCHandler(sessionService)
	if err != nil {
		log.Fatal(err)
	}

	grpcTransport, err := transport.NewSessionGRPCTransportImpl(grpcHandler)
	if err != nil {
		log.Fatal(err)
	}

	httpHandler, err := handler.NewSessionHTTPHandlerImpl(sessionService)
	if err != nil {
		log.Fatal(err)
	}
	httpTransport, err := transport.NewHTTPTransportImpl(httpHandler)
	if err != nil {
		log.Fatal(err)
	}
	httpConfig, err := config.LoadHTTPConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	grpcConfig, err := config.LoadGRPCConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Fatal(httpTransport.Listen(":" + httpConfig.Port))
	}()

	log.Fatal(grpcTransport.Listen(":" + grpcConfig.Port))
}
