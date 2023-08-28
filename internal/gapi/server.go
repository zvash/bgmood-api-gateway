package gapi

import (
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"github.com/zvash/bgmood-api-gateway/internal/util"
)

// Server serves gRPC requests for our api gateway.
type Server struct {
	pb.UnimplementedAppServer
	config util.Config
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}
