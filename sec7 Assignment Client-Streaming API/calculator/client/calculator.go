package main

import (
	"context"
	"log"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

func doCalculate(c pb.CalculatorClient) {
	log.Println("doCalculate was invoked")

	reqs := []int32{1, 2, 3, 4, 5}

	stream, err := c.Calculate(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Calculate %v\n", err)
	}

	for _, req := range reqs {

		stream.Send(&pb.CalculatorRequest{
			Num: req,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving Calculate %v\n", err)
	}

	log.Printf("Response from server %v\n", res.Average)

}
