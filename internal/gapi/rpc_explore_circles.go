package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
)

func (server *Server) ExploreCircles(ctx context.Context, req *cpb.ExploreCirclesRequest) (*cpb.ExploreCirclesResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.ExploreCircles(outgoingContext, req)
}
