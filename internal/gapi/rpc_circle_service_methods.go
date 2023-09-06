package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
)

func (server *Server) EditCircle(ctx context.Context, req *cpb.EditCircleRequest) (*cpb.EditCircleResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.EditCircle(outgoingContext, req)
}

func (server *Server) DestroyCircle(ctx context.Context, req *cpb.DestroyCircleRequest) (*cpb.DestroyCircleResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.DestroyCircle(outgoingContext, req)
}

func (server *Server) InviteUserToCircle(ctx context.Context, req *cpb.InviteUserToCircleRequest) (*cpb.InviteUserToCircleResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.InviteUserToCircle(outgoingContext, req)
}

func (server *Server) RollbackInviteUserToCircle(ctx context.Context, req *cpb.RollbackInviteUserToCircleRequest) (*cpb.RollbackInviteUserToCircleResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.RollbackInviteUserToCircle(outgoingContext, req)
}

func (server *Server) KickFromCircle(ctx context.Context, req *cpb.KickFromCircleRequest) (*cpb.KickFromCircleResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.KickFromCircle(outgoingContext, req)
}

func (server *Server) AcceptJoinRequest(ctx context.Context, req *cpb.AcceptJoinRequestRequest) (*cpb.AcceptJoinRequestResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.AcceptJoinRequest(outgoingContext, req)
}

func (server *Server) RemoveJoinRequest(ctx context.Context, req *cpb.RemoveJoinRequestRequest) (*cpb.RemoveJoinRequestResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.RemoveJoinRequest(outgoingContext, req)
}

func (server *Server) LeaveCircle(ctx context.Context, req *cpb.LeaveCircleRequest) (*cpb.LeaveCircleResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.LeaveCircle(outgoingContext, req)
}

func (server *Server) GetJoinedCircles(ctx context.Context, req *cpb.GetJoinedCirclesRequest) (*cpb.GetJoinedCirclesResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.GetJoinedCircles(outgoingContext, req)
}

func (server *Server) GetRequestedCircles(ctx context.Context, req *cpb.GetRequestedCirclesRequest) (*cpb.GetRequestedCirclesResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.GetRequestedCircles(outgoingContext, req)
}

func (server *Server) ViewCircle(ctx context.Context, req *cpb.ViewCircleRequest) (*cpb.ViewCircleResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.ViewCircle(outgoingContext, req)
}

func (server *Server) PromoteToPoster(ctx context.Context, req *cpb.PromoteToPosterRequest) (*cpb.AccessModificationResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.PromoteToPoster(outgoingContext, req)
}

func (server *Server) PromoteToAdmin(ctx context.Context, req *cpb.PromoteToAdminRequest) (*cpb.AccessModificationResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.PromoteToAdmin(outgoingContext, req)
}

func (server *Server) DemoteToViewer(ctx context.Context, req *cpb.DemoteToViewerRequest) (*cpb.AccessModificationResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.DemoteToViewer(outgoingContext, req)
}

func (server *Server) GetCircleWPChangeAccess(ctx context.Context, req *cpb.GetCircleWPChangeAccessRequest) (*cpb.CircleWPChangeAccessResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.GetCircleWPChangeAccess(outgoingContext, req)
}

func (server *Server) SetCircleWPChangeAccess(ctx context.Context, req *cpb.SetCircleWPChangeAccessRequest) (*cpb.CircleWPChangeAccessResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.SetCircleWPChangeAccess(outgoingContext, req)
}

func (server *Server) React(ctx context.Context, req *cpb.ReactRequest) (*cpb.MoodResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.React(outgoingContext, req)
}

func (server *Server) RemoveReaction(ctx context.Context, req *cpb.RemoveReactionRequest) (*cpb.MoodResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.RemoveReaction(outgoingContext, req)
}

func (server *Server) GetAvailableReactions(ctx context.Context, req *cpb.GetAvailableReactionsRequest) (*cpb.GetAvailableReactionsResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.GetAvailableReactions(outgoingContext, req)
}

func (server *Server) SendJoinRequest(ctx context.Context, req *cpb.SendJoinRequestRequest) (*cpb.SendJoinRequestResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}
	return server.CircleServiceClient.SendJoinRequest(outgoingContext, req)
}
