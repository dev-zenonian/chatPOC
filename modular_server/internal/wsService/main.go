package main

import (
	"common/client"
	"common/config"
	"log"
	"wsService/handler"
	"wsService/service"
	"wsService/transport"
)

func main() {
	// kafkaCfg, err := config.LoadKafkaConfig("./application/env/.env")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// kafkaConn, err := utils.InitKafka(kafkaCfg.Address, kafkaCfg.Topic, kafkaCfg.Partition)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	grpcClientCfg, err := config.LoadGRPCClientConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	sessionClient, err := client.NewSessionClientImpl(grpcClientCfg)
	if err != nil {
		log.Fatal(err)
	}

	// HARDCODE
	grpcClientCfg.Endpoint = "127.0.0.1:8090"
	messageClient, err := client.NewMessageClientImpl(grpcClientCfg)
	if err != nil {
		log.Fatal(err)
	}

	wsService, err := service.NewWSServiceImpl(sessionClient, messageClient)
	if err != nil {
		log.Fatal(err)
	}
	wsHTTPHandler, err := handler.NewWSHTTPHandlerImpl(wsService)
	if err != nil {
		log.Fatal(err)
	}
	wsGRPCHandler, err := handler.NewWSGRPCHandlerImpl(wsService)
	if err != nil {
		log.Fatal(err)
	}

	httpTransport, err := transport.NewHTTPTransportImpl(wsHTTPHandler)
	if err != nil {
		log.Fatal(err)
	}

	grpcTransport, err := transport.NewGRPCTransportImpl(wsGRPCHandler)
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
	go func() { log.Fatal(grpcTransport.Listen(":" + grpcConfig.Port)) }()
	log.Fatal(httpTransport.Listen(":" + httpConfig.Port))
}
