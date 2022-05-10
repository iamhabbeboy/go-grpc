package main

import (
	"context"
	pb "grpc_tutorial/gen/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := pb.NewTestApiClient(con)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello from grpc"})
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
	log.Println(resp.Msg)
}
