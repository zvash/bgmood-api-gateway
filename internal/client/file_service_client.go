package client

import (
	"context"
	"fmt"
	"github.com/zvash/bgmood-api-gateway/internal/filepb"
	"github.com/zvash/bgmood-api-gateway/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

type FileServiceClient struct {
	service filepb.FileClient
}

func NewFileServiceClient(cc *grpc.ClientConn) *FileServiceClient {
	service := filepb.NewFileClient(cc)
	return &FileServiceClient{
		service: service,
	}
}

func (client *FileServiceClient) UploadAvatar(
	inputStream *pb.App_UpdateProfileServer,
	ext string,
) (*filepb.UploadFileResponse, error) {
	iStream := *inputStream
	ctx := iStream.Context()
	outputStream, err := client.uploadFile(ctx, filepb.FileInfo_AVATAR_IMAGE, ext, nil)
	if err != nil {
		return nil, logError(status.Errorf(codes.Unknown, "couldn't initiate a stream to file service: %v", err))
	}

	imageSize := 0
	for {
		err := contextError(iStream.Context())
		if err != nil {
			return nil, logError(err)
		}
		req, err := iStream.Recv()

		var finished bool
		var chunk []byte
		if err == io.EOF {
			log.Print("upload is finished")
			finished = true
		}
		if err != nil && !finished {
			return nil, logError(status.Errorf(codes.Unknown, "cannot receive data from request stream: %v", err))
		}
		if !finished {
			chunk = req.GetChunkData()
			size := len(chunk)
			log.Printf("received a chunk with size: %d", size)
			imageSize += size
			log.Printf("total recieved: %d", imageSize)
			outputStream, err = client.uploadFile(ctx, filepb.FileInfo_AVATAR_IMAGE, ext, chunk, outputStream)
			if err == io.EOF {
				log.Print("upload to file service is finished")
				break
			}
			if err != nil {
				return nil, logError(status.Errorf(codes.Unknown, "failed to send data to file service: %v", err))
			}
		} else {
			break
		}
	}
	uploadResponse, oErr := (*outputStream).CloseAndRecv()
	err = iStream.SendAndClose(&pb.UpdateProfileResponse{
		Message: fmt.Sprintf("user was updated"),
	})
	if oErr != nil {
		return nil, logError(status.Errorf(codes.Unknown, "couldn't get the response from file server: %v", oErr))
	}
	if err != nil {
		return uploadResponse, logError(status.Errorf(codes.Unknown, "couldn't close the input stream: %v", oErr))
	}
	return uploadResponse, nil
}

func (client *FileServiceClient) uploadFile(
	ctx context.Context,
	ft filepb.FileInfo_FileType,
	ext string,
	chunk []byte,
	streams ...*filepb.File_UploadFileClient,
) (*filepb.File_UploadFileClient, error) {
	var stream filepb.File_UploadFileClient
	var req *filepb.UploadFileRequest
	var err error
	if len(streams) == 0 {
		stream, err = client.service.UploadFile(ctx)
		if err != nil {
			return nil, err
		}
		req = &filepb.UploadFileRequest{
			Data: &filepb.UploadFileRequest_Info{
				Info: &filepb.FileInfo{
					FileType:  ft,
					Extension: &ext,
				},
			},
		}
	} else {
		stream = *streams[0]
		req = &filepb.UploadFileRequest{
			Data: &filepb.UploadFileRequest_ChunkData{
				ChunkData: chunk,
			},
		}
	}
	err = stream.Send(req)
	if err != nil {
		return &stream, err
	}
	return &stream, nil
}
