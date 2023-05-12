package main

import (
	"io"
	"log"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

func (s *Server) Calculate(stream pb.Calculator_CalculateServer) error {
	log.Println("Calculate was invoked")

	res := 0

	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.CalculatorResponse{
				Average: int32(res) / int32(count),
			})
		}

		if err != nil {
			log.Fatalf("Something wrong %v\n", err)
		}

		res += int(req.Num)
		count++

	}

}
