package transport

import (
	"common/pb"
	"common/transport"
	"log"
	"messageService/handler"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcTransportImpl struct {
	messageHandler *handler.MessageGRPCHandler
	grpcServer     *grpc.Server
}

func NewMessageGRPCTransportImpl(messsageHandler *handler.MessageGRPCHandler) (transport.GRPCTransport, error) {
	tsp := &grpcTransportImpl{
		messageHandler: messsageHandler,
	}
	tsp.initConnection()
	return tsp, nil
}

func (h *grpcTransportImpl) initConnection() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterMessageServiceServer(grpcServer, h.messageHandler)
	h.grpcServer = grpcServer
}

func (h *grpcTransportImpl) Listen(port string) error {
	listenr, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	log.Printf("GRPC Listening on port %v\n", port)
	return h.grpcServer.Serve(listenr)
}
