package gapi

import (
	"fmt"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"github.com/zvash/bgmood-api-gateway/internal/util"
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
		var mood = &pb.Mood{}
		mood, err = util.ConvertTypes(circleServiceResp.Mood, mood)
		if err != nil {
			return logError(fmt.Errorf("could not convert circle-service circle to gateway circle: %w", err))
		}
		shareMoodResponse.Mood = mood
	}
	return nil
}
