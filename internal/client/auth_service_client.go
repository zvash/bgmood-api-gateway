package client

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
	"google.golang.org/grpc"
)

type AuthServiceClient struct {
	service authpb.AuthClient
}

func NewAuthServiceClient(cc *grpc.ClientConn) *AuthServiceClient {
	service := authpb.NewAuthClient(cc)
	return &AuthServiceClient{
		service: service,
	}
}

func (client *AuthServiceClient) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	resp, err := client.service.Login(ctx, req)
	return resp, err
}

func (client *AuthServiceClient) ListActiveSessions(ctx context.Context, req *authpb.ListActiveSessionsRequest) (*authpb.ListActiveSessionsResponse, error) {
	return client.service.ListActiveSessions(ctx, req)
}

func (client *AuthServiceClient) RegisterUser(ctx context.Context, req *authpb.RegisterUserRequest) (*authpb.RegisterUserResponse, error) {
	return client.service.RegisterUser(ctx, req)
}

func (client *AuthServiceClient) UpdateUser(ctx context.Context, req *authpb.UpdateUserRequest) (*authpb.UpdateUserResponse, error) {
	return client.service.UpdateUser(ctx, req)
}

func (client *AuthServiceClient) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	return client.service.Authenticate(ctx, req)
}

func (client *AuthServiceClient) GetUsersInfo(ctx context.Context, req *authpb.GetUsersInfoRequest) (*authpb.GetUsersInfoResponse, error) {
	return client.service.GetUsersInfo(ctx, req)
}
