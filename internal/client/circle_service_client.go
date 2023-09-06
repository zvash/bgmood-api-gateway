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

func (client *CircleServiceClient) CreateMood(
	ctx context.Context,
	req *cpb.CreateMoodRequest,
) (*cpb.MoodResponse, error) {
	return client.service.CreateMood(ctx, req)
}

func (client *CircleServiceClient) ExploreCircles(
	ctx context.Context,
	req *cpb.ExploreCirclesRequest,
) (*cpb.ExploreCirclesResponse, error) {
	return client.service.ExploreCircles(ctx, req)
}
