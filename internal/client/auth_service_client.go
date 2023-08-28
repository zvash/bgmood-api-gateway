package client

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
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
func (client *AuthServiceClient) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	authRequest := &authpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	authResponse, err := client.service.Login(ctx, authRequest)
	if err != nil {
		return nil, err
	}
	resp := &pb.LoginResponse{
		SessionId:             authResponse.SessionId,
		AccessToken:           authResponse.AccessToken,
		AccessTokenExpiresAt:  authResponse.AccessTokenExpiresAt,
		RefreshToken:          authResponse.RefreshToken,
		RefreshTokenExpiresAt: authResponse.RefreshTokenExpiresAt,
		Verified:              authResponse.Verified,
		User: &pb.User{
			Id:         authResponse.User.Id,
			Email:      authResponse.User.Email,
			Name:       authResponse.User.Name,
			Avatar:     authResponse.User.Avatar,
			IsVerified: authResponse.User.IsVerified,
			CreateAt:   authResponse.User.CreateAt,
		},
	}
	return resp, nil
}
