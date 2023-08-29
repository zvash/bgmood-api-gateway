package gapi

import (
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
	"github.com/zvash/bgmood-api-gateway/internal/client"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"github.com/zvash/bgmood-api-gateway/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Server serves gRPC requests for our api gateway.
type Server struct {
	pb.UnimplementedAppServer
	authpb.UnimplementedAuthServer
	AuthServiceClient *client.AuthServiceClient
	config            util.Config
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}
	serverAddress := server.config.AuthServiceGRPCServerAddress
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(serverAddress, transportOption)
	if err != nil {
		return nil, err
	}
	server.AuthServiceClient = client.NewAuthServiceClient(cc)
	return server, nil
}
