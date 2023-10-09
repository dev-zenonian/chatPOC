package main

import (
	"common/config"
	"log"
	"userService/handler"
	"userService/repository"
	"userService/service"
	"userService/transport"
)

func main() {
	repoCfg, err := config.LoadMongoConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := repository.NewUserRepositoryImpl(repoCfg)
	if err != nil {
		log.Fatal(err)
	}

	userService, err := service.NewUserServiceImpl(userRepo)
	if err != nil {
		log.Fatal(err)
	}

	grpcHandler, err := handler.NewUserGRPCHandler(userService)
	if err != nil {
		log.Fatal(err)
	}
	grpcTransport, err := transport.NewGRPCTransportImpl(grpcHandler)
	if err != nil {
		log.Fatal(err)
	}
	httpHandler, err := handler.NewUserHTTPHandler(userService)
	if err != nil {
		log.Fatal(err)
	}
	httpTransport, err := transport.NewHTTPTransportImpl(httpHandler)
	if err != nil {
		log.Fatal(err)
	}

	grpcCfg, err := config.LoadGRPCConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}

	httpCfg, err := config.LoadHTTPConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(grpcTransport.Listen(":" + grpcCfg.Port))
	}()

	log.Fatal(httpTransport.Listen(":" + httpCfg.Port))
	//
	// mux, err := transport.NewMuxTransport(grpcHandler, ":"+httpCfg.Port, ":"+grpcCfg.Port)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Fatal(mux.Listen())
}
