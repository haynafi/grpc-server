package main

import (
	"log"
	"net"

	"github.com/grpc-server/pb"
	"google.golang.org/grpc"
)

func main() {
	port := "1531"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterEmployeeServer(s, &employee.Server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
