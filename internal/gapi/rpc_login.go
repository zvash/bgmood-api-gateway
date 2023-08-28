package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/client"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	requestContext := server.buildClientContext(ctx)
	serverAddress := server.config.AuthServiceGRPCServerAddress
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(serverAddress, transportOption)
	if err != nil {
		return nil, err
	}
	authClient := client.NewAuthServiceClient(cc)
	resp, err := authClient.Login(requestContext, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
