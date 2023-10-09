package main

import (
	"common/client"
	"common/config"
	"groupService/handler"
	"groupService/repository"
	"groupService/service"
	transport "groupService/tranport"
	"log"
)

func main() {
	// kafkaCfg, err := config.LoadKafkaConfig("/application/env/.env") if err != nil { log.Fatal(err)
	// }
	grpcClientConfig, err := config.LoadGRPCClientConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	userClient, err := client.NewUserClientImpl(grpcClientConfig)
	mongoCfg, err := config.LoadMongoConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	groupRepository, err := repository.NewGroupRepositoryImpl(mongoCfg)
	if err != nil {
		log.Fatal(err)
	}

	sv, err := service.NewGroupServiceImpl(groupRepository, userClient)
	if err != nil {
		log.Fatal(err)
	}
	httpHandler, err := handler.NewGroupHTTPHandlerImpl(sv)
	if err != nil {
		log.Fatal(err)
	}

	grpcHandler, err := handler.NewGroupGRPCHandler(sv)
	if err != nil {
		log.Fatal(err)
	}

	httpTransport, err := transport.NewHTTPTransportImpl(httpHandler)
	if err != nil {
		log.Fatal(err)
	}

	grpcTransport, err := transport.NewGroupGRPCTransportImpl(grpcHandler)
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

	// go func() { log.Fatal(grpcTransport.Listen(":" + grpcCfg.Port)) }()
	// mux, err := transport.NewMuxTransport(grpcHandler, ":"+httpCfg.Port, ":"+grpcCfg.Port)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Fatal(mux.Listen())
}
