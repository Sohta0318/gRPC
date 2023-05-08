package main

import (
	"log"
	"net"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct{
pb.CalculatorServer
}

type test []string

func main(){
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Something wrong connection: %v\n", err)
	}

	log.Printf("Listening on: %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Something wrong connection: %v\n", err)
	}
}