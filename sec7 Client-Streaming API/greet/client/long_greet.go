package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Sohta0318/gRPC/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient){
log.Println("doLongGreet was invoked")

reqs := []*pb.GreetRequest{
	{FirstName: "Sohta"},
	{FirstName: "Marie"},
	{FirstName: "Test"},
}

stream, err := c.LongGreet(context.Background())

if err != nil {
	log.Fatalf("Error while calling LogGreet %v\n", err)
}

//* Sending request to server
for _, req := range reqs{
log.Printf("Sending req %v\n", req)

stream.Send(req)

time.Sleep(1 * time.Second)
}

//* Getting response from server

res, err := stream.CloseAndRecv()

if err != nil{
	log.Fatalf("Error while receiving response from LongGreet %v\n", err)
}

log.Printf("LongGreet: %v\n", res.Result)
}