package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
)

func (server *Server) SendJoinRequest(ctx context.Context, req *cpb.SendJoinRequestRequest) (*cpb.SendJoinRequestResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.SendJoinRequest(outgoingContext, req)
}
