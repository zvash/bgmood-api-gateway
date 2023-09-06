package client

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
	"google.golang.org/grpc"
)

type CircleServiceClient struct {
	service cpb.CirclesClient
}

func NewCircleServiceClient(cc *grpc.ClientConn) *CircleServiceClient {
	service := cpb.NewCirclesClient(cc)
	return &CircleServiceClient{
		service: service,
	}
}

func (client *CircleServiceClient) CreateCircle(
	ctx context.Context,
	req *cpb.CreateCircleRequest,
) (*cpb.CreateCircleResponse, error) {
	return client.service.CreateCircle(ctx, req)
}
