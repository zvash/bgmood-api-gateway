package main

import (
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
	"github.com/zvash/bgmood-api-gateway/internal/gapi"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"github.com/zvash/bgmood-api-gateway/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	runGrpcServer(config)
}

func runGrpcServer(config util.Config) {
	server, err := gapi.NewServer(config)
	if err != nil {
		log.Fatal("cannot create gRPC server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAppServer(grpcServer, server)
	authpb.RegisterAuthServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
