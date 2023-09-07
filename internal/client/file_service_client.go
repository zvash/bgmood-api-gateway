package client

import (
	"context"
	"github.com/zvash/bgmood-api-gateway/internal/filepb"
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

type ChunkDataAware interface {
	GetChunkData() []byte
}

func UploadFileWithStream(
	client *FileServiceClient,
	fileType filepb.FileInfo_FileType,
	ext string,
	req ChunkDataAware,
	inputStream grpc.ServerStream,
) (grpc.ServerStream, *filepb.UploadFileResponse, error) {
	ctx := inputStream.Context()
	outputStream, err := client.uploadFile(ctx, fileType, ext, nil)
	if err != nil {
		return inputStream, nil, logError(status.Errorf(codes.Unknown, "couldn't initiate a stream to file service: %v", err))
	}
	imageSize := 0
	for {
		err := contextError(inputStream.Context())
		if err != nil {
			return inputStream, nil, logError(err)
		}
		if err := inputStream.RecvMsg(req); err != nil {
			if err == io.EOF {
				log.Print("upload is finished")
			} else {
				return nil, nil, logError(status.Errorf(codes.Unknown, "cannot receive data from request stream: %v", err))
			}
			break
		}
		chunk := req.GetChunkData()
		size := len(chunk)
		log.Printf("received a chunk with size: %d", size)
		imageSize += size
		log.Printf("total received: %d", imageSize)

		outputStream, err = client.uploadFile(ctx, fileType, ext, chunk, outputStream)
		if err == io.EOF {
			log.Print("upload to file service is finished")
			break
		}
		if err != nil {
			return nil, nil, logError(status.Errorf(codes.Unknown, "failed to send data to file service: %v", err))
		}

	}

	uploadResponse, oErr := (*outputStream).CloseAndRecv()
	if oErr != nil {
		return nil, nil, logError(status.Errorf(codes.Unknown, "couldn't get the response from file server: %v", oErr))
	}
	return inputStream, uploadResponse, nil
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
