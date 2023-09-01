package gapi

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type HeaderPair struct {
	Key   string
	Value string
}

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	authorizationHeader        = "authorization"
	userObjectHeader           = "user-object"
)

func (server *Server) buildClientContext(ctx context.Context, headerPairs ...HeaderPair) context.Context {
	headers := make(map[string]string)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if authHeaders := md.Get(authorizationHeader); len(authHeaders) > 0 {
			headers[authorizationHeader] = authHeaders[0]
		}
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			headers[grpcGatewayUserAgentHeader] = userAgents[0]
		}
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			headers[userAgentHeader] = userAgents[0]
		}
		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			headers[xForwardedForHeader] = clientIPs[0]
		}
	}
	if p, ok := peer.FromContext(ctx); ok {
		headers[xForwardedForHeader] = p.Addr.String()
	}
	for _, pair := range headerPairs {
		headers[pair.Key] = pair.Value
	}
	newMetaData := metadata.New(headers)
	return metadata.NewOutgoingContext(ctx, newMetaData)
}
