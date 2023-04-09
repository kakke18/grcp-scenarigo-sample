package service

import (
	"context"

	"github.com/kakke18/grcp-scenarigo-sample/pb"
)

var _ pb.EchoServiceServer = (*EchoService)(nil)

type EchoService struct {
	pb.UnimplementedEchoServiceServer
}

func NewEchoService() *EchoService {
	return &EchoService{}
}

func (s *EchoService) Echo(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	message := req.GetMessage()

	return &pb.EchoResponse{
		Message: message,
	}, nil
}
