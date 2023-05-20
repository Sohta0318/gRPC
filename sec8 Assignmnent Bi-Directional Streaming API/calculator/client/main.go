package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Connecting server error %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	doSortNumber(c)
}
