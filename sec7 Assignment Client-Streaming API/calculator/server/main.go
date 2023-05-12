package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Connecting error", err)
	}
	log.Printf("Connecting :%v\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Something wrong grpc connecting")
	}

}
