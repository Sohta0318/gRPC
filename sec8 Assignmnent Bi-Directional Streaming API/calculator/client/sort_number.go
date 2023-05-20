package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Sohta0318/gRPC/calculator/proto"
)

func doSortNumber(c pb.CalculatorClient) {
	log.Println("doSortNumber was invoked")

	stream, err := c.SortNumber(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.CalculatorRequest{
				Num: number,
			})

			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Problem while reading server stream: %v\n", err)
				break
			}

			log.Printf("Received a new maximum of...: %v\n", res.Order)
		}
		close(waitc)
	}()
	<-waitc
}
