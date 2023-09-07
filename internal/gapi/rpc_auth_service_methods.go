package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
)

func (server *Server) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	return server.AuthServiceClient.RefreshToken(server.buildClientContext(ctx), req)
}

func (server *Server) TerminateSingleSession(ctx context.Context, req *authpb.TerminateSingleSessionRequest) (*authpb.TerminateSingleSessionResponse, error) {
	return server.AuthServiceClient.TerminateSingleSession(server.buildClientContext(ctx), req)
}

func (server *Server) TerminateOtherSessions(ctx context.Context, req *authpb.TerminateOtherSessionsRequest) (*authpb.TerminateOtherSessionsResponse, error) {
	return server.AuthServiceClient.TerminateOtherSessions(server.buildClientContext(ctx), req)
}

func (server *Server) Logout(ctx context.Context, req *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	return server.AuthServiceClient.Logout(server.buildClientContext(ctx), req)
}

func (server *Server) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	return server.AuthServiceClient.ChangePassword(server.buildClientContext(ctx), req)
}

func (server *Server) ResetPassword(ctx context.Context, req *authpb.ResetPasswordRequest) (*authpb.ResetPasswordResponse, error) {
	return server.AuthServiceClient.ResetPassword(server.buildClientContext(ctx), req)
}

func (server *Server) RequestPasswordReset(ctx context.Context, req *authpb.RequestPasswordResetRequest) (*authpb.RequestPasswordResetResponse, error) {
	return server.AuthServiceClient.RequestPasswordReset(server.buildClientContext(ctx), req)
}

func (server *Server) VerifyEmail(ctx context.Context, req *authpb.VerifyEmailRequest) (*authpb.VerifyEmailResponse, error) {
	return server.AuthServiceClient.VerifyEmail(server.buildClientContext(ctx), req)
}

func (server *Server) ResendVerificationEmail(ctx context.Context, req *authpb.ResendVerificationEmailRequest) (*authpb.ResendVerificationEmailResponse, error) {
	return server.AuthServiceClient.ResendVerificationEmail(server.buildClientContext(ctx), req)
}
