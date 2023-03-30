package main

import (
	"plugins/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn *grpc.ClientConn
)

func newConn(target string) *grpc.ClientConn {
	if conn != nil {
		return conn
	}

	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	return conn
}

func NewEchoServiceClient(address string) pb.EchoServiceClient {
	conn := newConn(address)

	return pb.NewEchoServiceClient(conn)
}

func NewUserServiceClient(address string) pb.UserServiceClient {
	conn := newConn(address)

	return pb.NewUserServiceClient(conn)
}
