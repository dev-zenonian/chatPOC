package main

import (
	"common/client"
	"common/config"
	"groupMessageService/handler"
	"groupMessageService/repository"
	"groupMessageService/service"
	"groupMessageService/transport"
	"log"
)

func main() {
	// kafkaCfg, err := config.LoadKafkaConfig("/application/env/.env") if err != nil { log.Fatal(err)
	// }
	grpcClientConfig, err := config.LoadGRPCClientConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	groupClient, err := client.NewGroupClient(grpcClientConfig)
	if err != nil {
		log.Fatal(err)
	}
	grpcClientConfig.Endpoint = "127.0.0.1:8088"
	distributorClient, err := client.NewDistributorClientImpl(grpcClientConfig)
	mongoCfg, err := config.LoadMongoConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	groupMessageRepository, err := repository.NewGroupMessageRepositoryImpl(mongoCfg)
	if err != nil {
		log.Fatal(err)
	}

	sv, err := service.NewGroupMessageServiceImpl(groupMessageRepository, groupClient, distributorClient)
	if err != nil {
		log.Fatal(err)
	}
	httpHandler, err := handler.NewGroupMessageHTTPHandler(sv)
	if err != nil {
		log.Fatal(err)
	}

	grpcHandler, err := handler.NewGroupMessageGRPCHandler(sv)
	if err != nil {
		log.Fatal(err)
	}

	httpTransport, err := transport.NewHTTPTransportImpl(httpHandler)
	if err != nil {
		log.Fatal(err)
	}

	grpcTransport, err := transport.NewSessionGRPCTransportImpl(grpcHandler)
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
}
