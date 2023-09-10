package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
	"github.com/zvash/bgmood-api-gateway/internal/gapi"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"github.com/zvash/bgmood-api-gateway/internal/prom"
	"github.com/zvash/bgmood-api-gateway/internal/util"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
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

	srvMetrics := prom.NewServerMetrics()
	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)
	exemplarFromContext := func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}
		return nil
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(prom.WithExemplarFromContext(exemplarFromContext)),
		),
		grpc.StreamInterceptor(
			srvMetrics.StreamServerInterceptor(prom.WithExemplarFromContext(exemplarFromContext)),
		),
	)

	pb.RegisterAppServer(grpcServer, server)
	authpb.RegisterAuthServer(grpcServer, server)
	cpb.RegisterCirclesServer(grpcServer, server)
	reflection.Register(grpcServer)

	srvMetrics.InitializeMetrics(grpcServer)

	go runPrometheusHttpServer(config, reg)

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

func runPrometheusHttpServer(config util.Config, reg *prometheus.Registry) {
	http.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics e.g. to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	err := http.ListenAndServe(config.PrometheusHTTPServerAddress, nil)
	if err != nil {
		log.Fatal("cannot start prometheus HTTP server")
	}
}
