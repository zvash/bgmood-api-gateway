package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
)

func (server *Server) ListActiveSessions(
	ctx context.Context,
	req *authpb.ListActiveSessionsRequest,
) (*authpb.ListActiveSessionsResponse, error) {
	requestContext := server.buildClientContext(ctx)
	return server.AuthServiceClient.ListActiveSessions(requestContext, req)
}
