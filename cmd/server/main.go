package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kakke18/grcp-scenarigo-sample/internal/service"
	"github.com/kakke18/grcp-scenarigo-sample/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = 8080
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, service.NewEchoService())

	reflection.Register(server)

	go func() {
		log.Printf("listening on port %d\n", port)
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	<-signalChan
	server.GracefulStop()
	log.Println("server exited")
}
