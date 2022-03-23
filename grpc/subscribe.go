package main

import (
	"context"
	"go-advanced/grpc/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func subscribe(conn *grpc.ClientConn) {
	client := proto.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &proto.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Println(reply.GetValue())
	}
}
