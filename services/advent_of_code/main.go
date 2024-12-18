package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	day_v1 "github.com/westford14/advent_of_code_2024/pkg/gen/westford14/advent_of_code/v1"
	rpcsrv "github.com/westford14/advent_of_code_2024/services/advent_of_code/internal/rpcsrv"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	server := rpcsrv.AdventServer{}
	day_v1.RegisterAdventServiceServer(s, &server)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
