package client

import (
	"common/config"
	"common/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGroupClient(cfg *config.GRPCClientConfig) (pb.GroupServiceClient, error) {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc2, err := grpc.Dial(
		cfg.Endpoint,
		transportOption,
	)
	if err != nil {
		return nil, err
	}
	client := pb.NewGroupServiceClient(cc2)
	return client, nil
}
