package gapi

import (
	"fmt"
	"github.com/zvash/bgmood-api-gateway/internal/cpb"
	"github.com/zvash/bgmood-api-gateway/internal/filepb"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"github.com/zvash/bgmood-api-gateway/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) BuildCircle(stream pb.App_BuildCircleServer) error {
	var buildCircleResponse = &pb.BuildCircleResponse{}
	defer func() {
		err := stream.SendAndClose(buildCircleResponse)
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
	cpbCreateCircleRequest := &cpb.CreateCircleRequest{
		Title:       req.GetInfo().GetTitle(),
		Description: req.GetInfo().GetDescription(),
		CircleType:  req.GetInfo().GetCircleType(),
		IsPrivate:   req.GetInfo().GetIsPrivate(),
	}

	extension := req.GetInfo().GetImageExt()
	buildCircleRequest := new(pb.BuildCircleRequest)
	_, uploadResponse, oErr := server.FileServiceClient.UploadFileWithStream(filepb.FileInfo_AVATAR_IMAGE, extension, buildCircleRequest, stream)
	if oErr != nil {
		return logError(oErr)
	}

	if uploadResponse != nil {
		cpbCreateCircleRequest.Avatar = uploadResponse.Path
		circleServiceResp, err := server.CircleServiceClient.CreateCircle(requestContext, cpbCreateCircleRequest)
		if err != nil {
			return logError(err)
		}
		var circle = &pb.Circle{}
		circle, err = util.ConvertTypes(circleServiceResp.Circle, circle)
		if err != nil {
			return logError(fmt.Errorf("could not convert circle-service circle to gateway circle: %w", err))
		}
		buildCircleResponse.Circle = circle
	}
	return nil
}
