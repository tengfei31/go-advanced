package main

import (
	"context"
	"go-advanced/grpc/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func channel(client proto.HelloServiceClient) {
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			reply := &proto.String{Value: "hi"}
			err = stream.Send(reply)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
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
	}()
}

func publish(conn *grpc.ClientConn) {
	client := proto.NewPubsubServiceClient(conn)

	_, err := client.Publish(context.Background(), &proto.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &proto.String{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}
