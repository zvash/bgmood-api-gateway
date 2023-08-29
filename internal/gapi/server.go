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
	FileServiceClient *client.FileServiceClient
	config            util.Config
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	authCC, err := dial(server.config.AuthServiceGRPCServerAddress)
	if err != nil {
		return nil, err
	}
	server.AuthServiceClient = client.NewAuthServiceClient(authCC)

	fileCC, err := dial(server.config.FileServiceGRPCServerAddress)
	if err != nil {
		return nil, err
	}
	server.FileServiceClient = client.NewFileServiceClient(fileCC)

	return server, nil
}

func dial(serverAddress string) (*grpc.ClientConn, error) {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(serverAddress, transportOption)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
