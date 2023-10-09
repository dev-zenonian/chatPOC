package transport

import (
	"common/pb"
	"context"
	"log"
	"messageService/handler"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MuxTransport struct {
	gHandler     *handler.MessageGRPCHandler
	mux          *runtime.ServeMux
	httpEndpoint string
	grpcEndpoint string
}

func NewMuxTransport(ghandler *handler.MessageGRPCHandler, httpEndpoint string, grpcEndpoint string) (*MuxTransport, error) {
	tsp := &MuxTransport{
		gHandler:     ghandler,
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
	err := pb.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, t.grpcEndpoint, opts)
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	t.mux = mux
	log.Println("Listen")
	return http.ListenAndServe(t.httpEndpoint, t.mux)

}
