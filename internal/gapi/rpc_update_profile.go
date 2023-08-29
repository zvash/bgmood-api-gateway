package gapi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/zvash/bgmood-api-gateway/internal/authpb"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
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

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		err := contextError(stream.Context())
		if err != nil {
			return logError(err)
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("upload is finished")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		imageSize += size
		log.Printf("total recieved: %d", imageSize)

		//send chunk to file service

		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	imagePath := fmt.Sprintf("/tmp/uploaded%s", extension)

	file, err := os.Create(imagePath)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "co: %v", err))
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot write image to file: %v", err))
	}

	err = stream.SendAndClose(&pb.UpdateProfileResponse{
		Message: fmt.Sprintf("user was updated"),
	})
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "could not close the stream"))
	}
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
