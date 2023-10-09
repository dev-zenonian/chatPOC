package main

import (
	"common/client"
	"common/config"
	"log"
	"messageService/handler"
	"messageService/repository"
	"messageService/service"
	"messageService/transport"
)

func main() {
	cfg, err := config.LoadMongoConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}

	messageRepo, err := repository.NewMessageRepositoryImpl(cfg)
	if err != nil {
		log.Fatal(err)
	}

	grpcClientCfg, err := config.LoadGRPCClientConfig("application/env/.env")
	userClient, err := client.NewUserClientImpl(grpcClientCfg)
	if err != nil {
		log.Fatal(err)
	}

	grpcClientCfg.Endpoint = "127.0.0.1:8088"
	distributorClient, err := client.NewDistributorClientImpl(grpcClientCfg)
	if err != nil {
		log.Fatal(err)
	}

	messageService, err := service.NewMessageServiceImpl(messageRepo, userClient, distributorClient)
	if err != nil {
		log.Fatal(err)
	}

	messageHTTPHandler, err := handler.NewMessageHTTPHandlerImpl(messageService)
	if err != nil {
		log.Fatal(err)
	}

	messageGRPCHandler, err := handler.NewMessageGRPCHandler(messageService)
	if err != nil {
		log.Fatal(err)
	}
	httpTransport, err := transport.NewMessageHTTPTransportImpl(messageHTTPHandler)
	if err != nil {
		log.Fatal(err)
	}
	grpcTransport, err := transport.NewMessageGRPCTransportImpl(messageGRPCHandler)
	if err != nil {
		log.Fatal(err)
	}
	httpCfg, err := config.LoadHTTPConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	grpcCfg, err := config.LoadGRPCConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}

	go func() { log.Fatal(httpTransport.Listen(":" + httpCfg.Port)) }()
	log.Fatal(grpcTransport.Listen(":" + grpcCfg.Port))

	// go func() {
	// 	log.Fatal(grpcTransport.Listen(":" + grpcCfg.Port))
	// }()
	// muxTransport, err := transport.NewMuxTransport(messageGRPCHandler, ":"+httpCfg.Port, ":"+grpcCfg.Port)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Fatal(muxTransport.Listen())
}
