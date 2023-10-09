package transport

import (
	"common/pb"
	"context"
	"log"
	"net/http"
	"userService/handler"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MuxTransport struct {
	ghandler     *handler.UserGRPCHandler
	mux          *runtime.ServeMux
	httpEndpoint string
	grpcEndpoint string
}

func NewMuxTransport(userHandler *handler.UserGRPCHandler, httpEndpoint string, grpcEndpoint string) (*MuxTransport, error) {
	tsp := &MuxTransport{
		ghandler:     userHandler,
		httpEndpoint: httpEndpoint,
		grpcEndpoint: grpcEndpoint,
	}
	return tsp, nil
}

func (t *MuxTransport) Listen() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Println(t.grpcEndpoint, t.httpEndpoint)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, t.grpcEndpoint, opts)
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	t.mux = mux
	log.Println("Listen")
	return http.ListenAndServe(t.httpEndpoint, t.mux)
}
