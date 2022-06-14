package main

import (
	pb "github.com/dtc03012/me/protobuf/proto/service"
	"github.com/dtc03012/me/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

const portNumber = "9001"

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMeServer(grpcServer, &service.MeServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
