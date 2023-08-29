package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
)

func (server *Server) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	requestContext := server.buildClientContext(ctx)
	resp, err := server.AuthServiceClient.Login(requestContext, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
