package client

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
	"google.golang.org/grpc"
)

type CircleServiceClient struct {
	service cpb.CirclesClient
}

func NewCircleServiceClient(cc *grpc.ClientConn) *CircleServiceClient {
	service := cpb.NewCirclesClient(cc)
	return &CircleServiceClient{
		service: service,
	}
}

func (client *CircleServiceClient) CreateCircle(
	ctx context.Context,
	req *cpb.CreateCircleRequest,
) (*cpb.CreateCircleResponse, error) {
	return client.service.CreateCircle(ctx, req)
}

func (client *CircleServiceClient) CreateMood(
	ctx context.Context,
	req *cpb.CreateMoodRequest,
) (*cpb.MoodResponse, error) {
	return client.service.CreateMood(ctx, req)
}

func (client *CircleServiceClient) ExploreCircles(
	ctx context.Context,
	req *cpb.ExploreCirclesRequest,
) (*cpb.ExploreCirclesResponse, error) {
	return client.service.ExploreCircles(ctx, req)
}

func (client *CircleServiceClient) SendJoinRequest(ctx context.Context, req *cpb.SendJoinRequestRequest) (*cpb.SendJoinRequestResponse, error) {
	return client.service.SendJoinRequest(ctx, req)
}

func (client *CircleServiceClient) EditCircle(ctx context.Context, req *cpb.EditCircleRequest) (*cpb.EditCircleResponse, error) {
	return client.service.EditCircle(ctx, req)
}

func (client *CircleServiceClient) DestroyCircle(ctx context.Context, req *cpb.DestroyCircleRequest) (*cpb.DestroyCircleResponse, error) {
	return client.service.DestroyCircle(ctx, req)
}

func (client *CircleServiceClient) InviteUserToCircle(ctx context.Context, req *cpb.InviteUserToCircleRequest) (*cpb.InviteUserToCircleResponse, error) {
	return client.service.InviteUserToCircle(ctx, req)
}

func (client *CircleServiceClient) RollbackInviteUserToCircle(ctx context.Context, req *cpb.RollbackInviteUserToCircleRequest) (*cpb.RollbackInviteUserToCircleResponse, error) {
	return client.service.RollbackInviteUserToCircle(ctx, req)
}

func (client *CircleServiceClient) KickFromCircle(ctx context.Context, req *cpb.KickFromCircleRequest) (*cpb.KickFromCircleResponse, error) {
	return client.service.KickFromCircle(ctx, req)
}

func (client *CircleServiceClient) AcceptJoinRequest(ctx context.Context, req *cpb.AcceptJoinRequestRequest) (*cpb.AcceptJoinRequestResponse, error) {
	return client.service.AcceptJoinRequest(ctx, req)
}

func (client *CircleServiceClient) RemoveJoinRequest(ctx context.Context, req *cpb.RemoveJoinRequestRequest) (*cpb.RemoveJoinRequestResponse, error) {
	return client.service.RemoveJoinRequest(ctx, req)
}

func (client *CircleServiceClient) LeaveCircle(ctx context.Context, req *cpb.LeaveCircleRequest) (*cpb.LeaveCircleResponse, error) {
	return client.service.LeaveCircle(ctx, req)
}

func (client *CircleServiceClient) GetJoinedCircles(ctx context.Context, req *cpb.GetJoinedCirclesRequest) (*cpb.GetJoinedCirclesResponse, error) {
	return client.service.GetJoinedCircles(ctx, req)
}

func (client *CircleServiceClient) GetRequestedCircles(ctx context.Context, req *cpb.GetRequestedCirclesRequest) (*cpb.GetRequestedCirclesResponse, error) {
	return client.service.GetRequestedCircles(ctx, req)
}

func (client *CircleServiceClient) ViewCircle(ctx context.Context, req *cpb.ViewCircleRequest) (*cpb.ViewCircleResponse, error) {
	return client.service.ViewCircle(ctx, req)
}

func (client *CircleServiceClient) PromoteToPoster(ctx context.Context, req *cpb.PromoteToPosterRequest) (*cpb.AccessModificationResponse, error) {
	return client.service.PromoteToPoster(ctx, req)
}

func (client *CircleServiceClient) PromoteToAdmin(ctx context.Context, req *cpb.PromoteToAdminRequest) (*cpb.AccessModificationResponse, error) {
	return client.service.PromoteToAdmin(ctx, req)
}

func (client *CircleServiceClient) DemoteToViewer(ctx context.Context, req *cpb.DemoteToViewerRequest) (*cpb.AccessModificationResponse, error) {
	return client.service.DemoteToViewer(ctx, req)
}

func (client *CircleServiceClient) GetCircleWPChangeAccess(ctx context.Context, req *cpb.GetCircleWPChangeAccessRequest) (*cpb.CircleWPChangeAccessResponse, error) {
	return client.service.GetCircleWPChangeAccess(ctx, req)
}

func (client *CircleServiceClient) SetCircleWPChangeAccess(ctx context.Context, req *cpb.SetCircleWPChangeAccessRequest) (*cpb.CircleWPChangeAccessResponse, error) {
	return client.service.SetCircleWPChangeAccess(ctx, req)
}

func (client *CircleServiceClient) React(ctx context.Context, req *cpb.ReactRequest) (*cpb.MoodResponse, error) {
	return client.service.React(ctx, req)
}

func (client *CircleServiceClient) RemoveReaction(ctx context.Context, req *cpb.RemoveReactionRequest) (*cpb.MoodResponse, error) {
	return client.service.RemoveReaction(ctx, req)
}

func (client *CircleServiceClient) GetAvailableReactions(ctx context.Context, req *cpb.GetAvailableReactionsRequest) (*cpb.GetAvailableReactionsResponse, error) {
	return client.service.GetAvailableReactions(ctx, req)
}

func (client *CircleServiceClient) GetCircleJoinRequests(ctx context.Context, req *cpb.GetCircleJoinRequestsRequest) (*cpb.GetCircleJoinRequestsResponse, error) {
	return client.service.GetCircleJoinRequests(ctx, req)
}
