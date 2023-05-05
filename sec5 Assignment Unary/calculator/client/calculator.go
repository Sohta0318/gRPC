package main

import (
	"context"
	"log"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

func doCalc(c pb.CalculatorClient) {

	log.Println("doCalc was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 10,
		Num2: 3,
	})

	if err != nil {
		log.Fatalf("Something wrong: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Sum)
}
