package transport

import (
	"common/pb"
	"common/transport"
	"groupService/handler"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcTransportImpl struct {
	groupHandler *handler.GroupGRPCHandler
	grpcServer   *grpc.Server
}

func NewGroupGRPCTransportImpl(groupHandler *handler.GroupGRPCHandler) (transport.GRPCTransport, error) {
	tsp := &grpcTransportImpl{
		groupHandler: groupHandler,
	}
	tsp.initConnection()
	return tsp, nil
}

func (h *grpcTransportImpl) initConnection() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterGroupServiceServer(grpcServer, h.groupHandler)
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
