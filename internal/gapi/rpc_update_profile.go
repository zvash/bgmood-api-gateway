package gapi

import (
	"fmt"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
	"github.com/zvash/bgmood-api-gateway/internal/client"
	"github.com/zvash/bgmood-api-gateway/internal/filepb"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (server *Server) UpdateProfile(stream pb.App_UpdateProfileServer) error {
	requestContext := server.buildClientContext(stream.Context())
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive the request.")
	}
	authUpdateRequest := &authpb.UpdateUserRequest{
		Name: req.GetInfo().Name,
	}
	_, err = server.AuthServiceClient.UpdateUser(requestContext, authUpdateRequest)
	if err != nil {
		return logError(err)
	}

	if req.GetInfo().GetImageExt() == "" {
		err = stream.SendAndClose(&pb.UpdateProfileResponse{
			Message: fmt.Sprintf("user was updated"),
		})
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "could not close the stream"))
		}
		return nil
	}
	extension := req.GetInfo().GetImageExt()

	updateProfileRequest := new(pb.UpdateProfileRequest)
	_, uploadResponse, oErr := client.UploadFileWithStream(server.FileServiceClient, filepb.FileInfo_AVATAR_IMAGE, extension, updateProfileRequest, stream)
	err = stream.SendAndClose(&pb.UpdateProfileResponse{
		Message: fmt.Sprintf("user was updated"),
	})
	if oErr != nil {
		return logError(oErr)
	}
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "couldn't close the input stream: %v", err))
	}
	if uploadResponse != nil {
		requestContext = server.buildClientContext(stream.Context())
		authUpdateRequest = &authpb.UpdateUserRequest{
			Avatar: &uploadResponse.Path,
		}
		_, err := server.AuthServiceClient.UpdateUser(requestContext, authUpdateRequest)
		if err != nil {
			return logError(err)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
