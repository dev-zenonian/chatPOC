package transport

import (
	"common/pb"
	"common/transport"
	"log"
	"net"
	"wsService/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcTransportImpl struct {
	wsHandler  *handler.WSGPRCHandler
	grpcServer *grpc.Server
}

func NewGRPCTransportImpl(wsHandler *handler.WSGPRCHandler) (transport.GRPCTransport, error) {
	tsp := &grpcTransportImpl{
		wsHandler: wsHandler,
	}
	tsp.initConnection()
	return tsp, nil
}

func (h *grpcTransportImpl) initConnection() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterWSHandlerServiceServer(grpcServer, h.wsHandler)
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
