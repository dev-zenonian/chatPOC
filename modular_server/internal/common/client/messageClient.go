package client

import (
	"common/config"
	"common/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMessageClientImpl(cfg *config.GRPCClientConfig) (pb.MessageServiceClient, error) {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc2, err := grpc.Dial(
		cfg.Endpoint,
		transportOption,
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewMessageServiceClient(cc2)
	return client, nil

}
