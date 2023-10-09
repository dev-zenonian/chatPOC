package transport

import (
	"common/pb"
	"common/transport"
	"groupMessageService/handler"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcTransportImpl struct {
	groupMessageHandler *handler.GroupMessageGRPCHandler
	grpcServer          *grpc.Server
}

func NewSessionGRPCTransportImpl(groupMessageHandler *handler.GroupMessageGRPCHandler) (transport.GRPCTransport, error) {
	tsp := &grpcTransportImpl{
		groupMessageHandler: groupMessageHandler,
	}
	tsp.initConnection()
	return tsp, nil
}

func (h *grpcTransportImpl) initConnection() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterGroupMessageServiceServer(grpcServer, h.groupMessageHandler)
	h.grpcServer = grpcServer
}

func (h *grpcTransportImpl) Listen(port string) error {
	listenr, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	log.Printf("Listening on port: %v", port)
	return h.grpcServer.Serve(listenr)
}
