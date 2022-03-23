package main

import (
	"crypto/tls"
	"crypto/x509"
	"go-advanced/grpc/proto"
	"io/ioutil"
	"log"
	"net"
	"testing"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TestHelloServiceImpl(t *testing.T) {
	// mux := http.NewServeMux()

	// http2Handler := h2c.NewHandler(mux, &http2.Server{})
	// server := &http.Server{Addr: ":3999", Handler: http2Handler}
	// server.ListenAndServe()

	// 利用第三方中间件来增加多个过滤器
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(filter)))
	proto.RegisterHelloServiceServer(grpcServer, NewHelloServiceImpl())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}

// CA根证书、server.crt
func TestPubsubServiceTLS(t *testing.T) {
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
func TestPubsubService(t *testing.T) {
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

func TestPubsubServiceNoTLS(t *testing.T) {
	grpcServer := grpc.NewServer()
	proto.RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}