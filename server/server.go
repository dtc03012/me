package main

import (
	"context"
	"github.com/dtc03012/me/handler"
	pb "github.com/dtc03012/me/protobuf/proto/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

const (
	grpcGatewayPortNumber = "9000"
	grpcPortNumber        = "9001"
	filePortNumber        = "4500"
)

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":"+grpcPortNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMeServer(grpcServer, handler.NewMeServer())

	log.Printf("start gRPC server on %s port", grpcPortNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func startGatewayServer(ctx context.Context) {
	mux := runtime.NewServeMux()
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if err := pb.RegisterMeHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+grpcPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("start HTTP server on %s port", grpcGatewayPortNumber)
	if err := http.ListenAndServe(":"+grpcGatewayPortNumber, mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func startServer(ctx context.Context) {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		startGRPCServer()
		wg.Done()
	}()

	go func() {
		startGatewayServer(ctx)
		wg.Done()
	}()

	wg.Wait()
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	startServer(ctx)
}
