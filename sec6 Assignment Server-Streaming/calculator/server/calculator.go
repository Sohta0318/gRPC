package main

import (
	"log"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

func (s *Server)Decompose(in *pb.CalculatorRequest, stream pb.Calculator_DecomposeServer) error{
	log.Printf("Decompose function was invoked with: %v\n", in)

	number := in.One
	divisor := int64(2)
	
	for number > 1 {
		if number % int32(divisor) == 0{
			stream.Send(&pb.CalculatorResponse{
				PrimeNumber: int32(divisor),
			})

			number /= int32(divisor)
		}else{
			divisor++
		}
	}

	return nil
}