package main

import (
	"plugins/echo/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewEchoServiceClient(address string) pb.EchoServiceClient {
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	return pb.NewEchoServiceClient(conn)
}
