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

func (client *AuthServiceClient) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	return client.service.RefreshToken(ctx, req)
}

func (client *AuthServiceClient) TerminateSingleSession(ctx context.Context, req *authpb.TerminateSingleSessionRequest) (*authpb.TerminateSingleSessionResponse, error) {
	return client.service.TerminateSingleSession(ctx, req)
}

func (client *AuthServiceClient) TerminateOtherSessions(ctx context.Context, req *authpb.TerminateOtherSessionsRequest) (*authpb.TerminateOtherSessionsResponse, error) {
	return client.service.TerminateOtherSessions(ctx, req)
}

func (client *AuthServiceClient) Logout(ctx context.Context, req *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	return client.service.Logout(ctx, req)
}

func (client *AuthServiceClient) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	return client.service.ChangePassword(ctx, req)
}

func (client *AuthServiceClient) ResetPassword(ctx context.Context, req *authpb.ResetPasswordRequest) (*authpb.ResetPasswordResponse, error) {
	return client.service.ResetPassword(ctx, req)
}

func (client *AuthServiceClient) RequestPasswordReset(ctx context.Context, req *authpb.RequestPasswordResetRequest) (*authpb.RequestPasswordResetResponse, error) {
	return client.service.RequestPasswordReset(ctx, req)
}

func (client *AuthServiceClient) VerifyEmail(ctx context.Context, req *authpb.VerifyEmailRequest) (*authpb.VerifyEmailResponse, error) {
	return client.service.VerifyEmail(ctx, req)
}

func (client *AuthServiceClient) ResendVerificationEmail(ctx context.Context, req *authpb.ResendVerificationEmailRequest) (*authpb.ResendVerificationEmailResponse, error) {
	return client.service.ResendVerificationEmail(ctx, req)
}
