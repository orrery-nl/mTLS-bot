package main

import (
	"google.golang.org/grpc"
	"log/slog"
	"mtls_bot_root/configuration"
	insecure_services_v1 "mtls_bot_root/services/v1/insecure"
	"net"
	"os"
	"sync"
)

var (
	config *configuration.RootConfiguration = configuration.Get()

	addressInsecure string = "localhost:50051"
	addressSecure   string = "localhost:50052"
)

func init() {
	if os.Getenv("MTLS_BOT_INSECURE_ADDRESS") != "" {
		addressInsecure = os.Getenv("MTLS_BOT_INSECURE_ADDRESS")
	}

	if os.Getenv("MTLS_BOT_SECURE_ADDRESS") != "" {
		addressSecure = os.Getenv("MTLS_BOT_SECURE_ADDRESS")
	}
}

func main() {

	waitGroup := sync.WaitGroup{}

	// Start the insecure gRPC server.
	//
	waitGroup.Add(1)
	go func() {
		insecureGrpcServer()
		waitGroup.Done()
	}()

	// Start the secure gRPC server.
	//
	waitGroup.Add(1)
	go func() {
		secureGrpcServer()
		waitGroup.Done()
	}()

	waitGroup.Wait()

	config.Store()
}

// insecureGrpcServer - Start the insecure gRPC server.
func insecureGrpcServer() {
	listener, err := net.Listen("tcp", addressInsecure)
	if err != nil {
		slog.Error("Failed to listen.", "error", err)
		os.Exit(1)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register services.
	//
	insecure_services_v1.RegisterAuthenticationServer(grpcServer, &insecure_services_v1.AuthenticationService{})

	grpcServer.Serve(listener)
}

// secureGrpcServer - Start the secure gRPC server.
func secureGrpcServer() {
	listener, err := net.Listen("tcp", addressSecure)
	if err != nil {
		slog.Error("Failed to listen.", "error", err)
		os.Exit(1)
	}

	// Register services.
	//

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpcServer.Serve(listener)
}
