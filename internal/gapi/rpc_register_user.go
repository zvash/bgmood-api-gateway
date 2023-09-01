package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
)

func (server *Server) RegisterUser(ctx context.Context, req *authpb.RegisterUserRequest) (*authpb.RegisterUserResponse, error) {
	return server.AuthServiceClient.RegisterUser(ctx, req)
}
