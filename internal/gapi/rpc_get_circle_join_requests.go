package gapi

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
)

func (server *Server) GetCircleJoinRequests(
	ctx context.Context,
	req *pb.GetCircleJoinRequestsRequest,
) (*pb.CircleJoinRequestsWithUserResponse, error) {
	outgoingContext, err := server.buildClientContextWithUser(ctx)
	if err != nil {
		return nil, err
	}

	circlesReq := &cpb.GetCircleJoinRequestsRequest{
		CircleId: req.GetCircleId(),
		Page:     req.GetPage(),
	}

	circleResp, err := server.CircleServiceClient.GetCircleJoinRequests(outgoingContext, circlesReq)
	if err != nil {
		return nil, err
	}

	userIds := make([]string, 0)
	userIdToJoinRequestMap := make(map[string]*cpb.JoinRequest)
	joinRequests := circleResp.GetJoinRequests()
	for _, joinRequest := range joinRequests {
		userIds = append(userIds, joinRequest.GetUserId())
		userIdToJoinRequestMap[joinRequest.GetUserId()] = joinRequest
	}

	authReq := &authpb.GetUsersInfoRequest{
		UserIds: userIds,
	}

	outgoingContext = server.buildClientContext(ctx)

	authResp, err := server.AuthServiceClient.GetUsersInfo(outgoingContext, authReq)
	if err != nil {
		return nil, err
	}
	usersInfo := authResp.GetUsers()
	userIdToUserInfoMap := make(map[string]*authpb.User)
	for _, userInfo := range usersInfo {
		userIdToUserInfoMap[userInfo.GetId()] = userInfo
	}

	usersJoinRequests := make([]*pb.UserJoinRequest, 0)
	for userId, joinRequestRecord := range userIdToJoinRequestMap {
		if userInfo, ok := userIdToUserInfoMap[userId]; ok {
			userJoinRequest := &pb.UserJoinRequest{
				UserId:        userId,
				Name:          userInfo.GetName(),
				Avatar:        userInfo.Avatar,
				JoinRequestId: joinRequestRecord.GetJoinRequestId(),
			}
			usersJoinRequests = append(usersJoinRequests, userJoinRequest)
		}
	}
	resp := &pb.CircleJoinRequestsWithUserResponse{
		Users: usersJoinRequests,
	}
	return resp, nil
}
