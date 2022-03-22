package main

import (
	"go-advanced/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	creds, err := credentials.NewServerTLSFromFile("secret/server.crt", "secret/server.key")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
