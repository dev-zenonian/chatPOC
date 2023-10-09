package transport

import (
	"common/pb"
	"common/transport"
	"distributorService/handler"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcTransportImpl struct {
	distributorHandler *handler.MessageDistributorGRPCHandler
	grpcServer         *grpc.Server
}

func NewGRPCTransportImpl(distributorHandler *handler.MessageDistributorGRPCHandler) (transport.GRPCTransport, error) {
	tsp := &grpcTransportImpl{
		distributorHandler: distributorHandler,
	}
	tsp.initConnection()
	return tsp, nil
}

func (h *grpcTransportImpl) initConnection() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterMessageDistributorServer(grpcServer, h.distributorHandler)
	h.grpcServer = grpcServer
}

func (h *grpcTransportImpl) Listen(port string) error {
	listenr, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	log.Printf("GRPC listening on %v\n", port)
	return h.grpcServer.Serve(listenr)
}
