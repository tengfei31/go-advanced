package main

import (
	"go-advanced/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	pubsubService()
}

func helloServiceImpl() {
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, NewHelloServiceImpl())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}

func pubsubService() {
	grpcServer := grpc.NewServer()
	proto.RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
