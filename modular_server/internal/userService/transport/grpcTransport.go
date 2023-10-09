package transport

import (
	"common/pb"
	"common/transport"
	"log"
	"net"
	"userService/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcTransportImpl struct {
	userHandler *handler.UserGRPCHandler
	grpcServer  *grpc.Server
}

func NewGRPCTransportImpl(userHandler *handler.UserGRPCHandler) (transport.GRPCTransport, error) {
	tsp := &grpcTransportImpl{
		userHandler: userHandler,
	}
	tsp.initConnection()
	return tsp, nil
}

func (h *grpcTransportImpl) initConnection() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterUserServiceServer(grpcServer, h.userHandler)
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
