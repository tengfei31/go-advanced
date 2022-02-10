package main

import (
	"context"
	"go-advanced/grpc/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "client"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.GetValue())

}
