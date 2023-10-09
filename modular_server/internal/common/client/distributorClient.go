package client

import (
	"common/config"
	"common/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewDistributorClientImpl(cfg *config.GRPCClientConfig) (pb.MessageDistributorClient, error) {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc2, err := grpc.Dial(
		cfg.Endpoint,
		transportOption,
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewMessageDistributorClient(cc2)
	return client, nil
}
