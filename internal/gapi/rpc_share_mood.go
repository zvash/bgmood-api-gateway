package gapi

import (
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ShareMood(stream pb.App_ShareMoodServer) error {
	var shareMoodResponse = &pb.ShareMoodResponse{}
	defer func() {
		err := stream.SendAndClose(shareMoodResponse)
		if err != nil {
			_ = logError(err)
		}
	}()
	requestContext, err := server.buildClientContextWithUser(stream.Context())
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "unauthorized.")
	}
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive the request.")
	}
	cpbCreateMoodRequest := &cpb.CreateMoodRequest{
		CircleId:      req.GetInfo().GetCircleId(),
		Description:   req.GetInfo().Description,
		SetBackground: req.GetInfo().GetSetBackground(),
	}

	extension := req.GetInfo().GetImageExt()
	_, uploadResponse, oErr := server.FileServiceClient.UploadMoodBackground(&stream, extension)
	if oErr != nil {
		return logError(oErr)
	}
	if uploadResponse != nil {
		cpbCreateMoodRequest.Image = uploadResponse.Path
		circleServiceResp, err := server.CircleServiceClient.CreateMood(requestContext, cpbCreateMoodRequest)
		if err != nil {
			return logError(err)
		}
		var mood = &pb.Mood{
			Id:            circleServiceResp.GetMood().GetId(),
			PosterId:      circleServiceResp.GetMood().GetPosterId(),
			CircleId:      circleServiceResp.GetMood().GetCircleId(),
			Image:         circleServiceResp.GetMood().GetImage(),
			SetBackground: circleServiceResp.GetMood().GetSetBackground(),
			CreatedAt:     circleServiceResp.GetMood().GetCreatedAt(),
			Description:   circleServiceResp.GetMood().Description,
		}
		shareMoodResponse.Mood = mood
	}
	return nil
}
