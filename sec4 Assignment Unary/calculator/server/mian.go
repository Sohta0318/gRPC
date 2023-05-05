package main

import (
	"log"
	"net"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Connect failed: %v\n", err)
	}

	log.Printf("Listening on: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Connect failed: %v\n", err)
	}
}
