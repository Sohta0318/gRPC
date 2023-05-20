package main

import (
	"io"
	"log"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

func (s *Server) SortNumber(stream pb.Calculator_SortNumberServer) error {
log.Println("sortNumber was invoked")

var max int32 = 0

for {
	req, err := stream.Recv()

	if err == io.EOF{
		return nil
	}

	if err != nil {
		log.Fatalf("Error while reading client stream: %v\n", err)
	}

	if number := req.Num; number > max{
		max = number

		err := stream.Send(&pb.CalculatorResponse{
			Order: max,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}



}
