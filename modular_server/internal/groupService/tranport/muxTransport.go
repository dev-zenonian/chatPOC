package transport

import (
	"common/pb"
	"context"
	"groupService/handler"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MuxTransport struct {
	mux          *runtime.ServeMux
	ghandler     *handler.GroupGRPCHandler
	httpEndpoint string
	grpcEndpoint string
}

func NewMuxTransport(ghandler *handler.GroupGRPCHandler, httpEndpoint string, grpcEndpoint string) (*MuxTransport, error) {
	mTransport := &MuxTransport{
		ghandler:     ghandler,
		httpEndpoint: httpEndpoint,
		grpcEndpoint: grpcEndpoint,
	}
	return mTransport, nil
}

func (t *MuxTransport) Listen() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Println(t.grpcEndpoint, t.httpEndpoint)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterGroupServiceHandlerFromEndpoint(ctx, mux, t.grpcEndpoint, opts)
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	t.mux = mux
	log.Println("Listn")
	return http.ListenAndServe(t.httpEndpoint, t.mux)
}
