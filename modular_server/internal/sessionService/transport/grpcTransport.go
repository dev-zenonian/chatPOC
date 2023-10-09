package transport

import (
	"common/pb"
	"common/transport"
	"log"
	"net"
	"sessionService/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcTransportImpl struct {
	sessionHandler *handler.SessionGRPCHandler
	grpcServer     *grpc.Server
}

func NewSessionGRPCTransportImpl(sessionHandler *handler.SessionGRPCHandler) (transport.GRPCTransport, error) {
	tsp := &grpcTransportImpl{
		sessionHandler: sessionHandler,
	}
	tsp.initConnection()
	return tsp, nil
}

func (h *grpcTransportImpl) initConnection() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterSessionServiceServer(grpcServer, h.sessionHandler)
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
