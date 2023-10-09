package main

import (
	"common/client"
	"common/config"
	"distributorService/handler"
	"distributorService/service"
	"distributorService/transport"
	"log"
)

func main() {
	grpcClientCfg, err := config.LoadGRPCClientConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	sessionClient, err := client.NewSessionClientImpl(grpcClientCfg)
	if err != nil {
		log.Fatal(err)
	}
	grpcConfig, err := config.LoadGRPCConfig("application/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	distributor, err := service.NewMessageDistributorImpl(sessionClient)
	if err != nil {
		log.Fatal(err)
	}
	grpcHandler, err := handler.NewMessageDistributoGRPCHandler(distributor)
	if err != nil {
		log.Fatal(err)
	}
	grpcTransport, err := transport.NewGRPCTransportImpl(grpcHandler)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(grpcTransport.Listen(":" + grpcConfig.Port))
}
