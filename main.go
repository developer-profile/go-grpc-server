package main

import (
	"fmt"
	"log"
	"net"

	"github.com/developer-profile/go-grpc-server/api/birthday"
	"google.golang.org/grpc"
)

func main() {
	// Listen on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Server instance to listen to
	s := birthday.Server{}

	grpcServer := grpc.NewServer()

	// attach service to server
	birthday.RegisterBirthdayCheckerServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve: %s", err)
	}

}
