package client

import (
	"common/config"
	"common/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewSessionClientImpl(cfg *config.GRPCClientConfig) (pb.SessionServiceClient, error) {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc2, err := grpc.Dial(
		cfg.Endpoint,
		transportOption,
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewSessionServiceClient(cc2)
	return client, nil
}
