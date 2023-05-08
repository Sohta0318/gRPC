package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

func doDecompose(c pb.CalculatorClient){
	log.Println("doDecompose was invoked")

	req := &pb.CalculatorRequest{
		One: 120,
	}

	stream, err:= c.Decompose(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for{
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Decompose: %d\n", msg.PrimeNumber)
	}
}