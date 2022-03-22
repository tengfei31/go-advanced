package main

import (
	"crypto/tls"
	"crypto/x509"
	"go-advanced/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//pubsubService()
	pubsubServiceTLS()
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

// CA根证书、server.crt
func pubsubServiceTLS() {
	certificate, err := tls.LoadX509KeyPair("secret/server.crt", "secret/server.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("secret/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}

// 根据server.crt进行校验
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
