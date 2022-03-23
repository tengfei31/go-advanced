package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"go-advanced/grpc/proto"
	"io/ioutil"
	"log"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// 利用客户端证书和CA根证书、服务器名字验证
func TestMainClientTLS(t *testing.T) {
	var tlsServerName string = "server.io"
	certificate, err := tls.LoadX509KeyPair("secret/client.crt", "secret/client.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("secret/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   tlsServerName,
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	publish(conn)
}

// 利用服务器的证书、服务器名字认证
func TestMainServerTLS(t *testing.T) {
	creds, err := credentials.NewClientTLSFromFile("secret/server.crt", "server.grpc.io")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	publish(conn)
}

// 无证书认证
func TestMainNoTLS(t *testing.T) {
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}
	conn, err := grpc.Dial(":1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	publish(conn)
}

// 基础的GRPC
func TestMainBaseGrpc(t *testing.T) {
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}
	conn, err := grpc.Dial(":1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	//channel(client)
	//for {
	//	time.Sleep(time.Hour)
	//}

	reply, err := client.Hello(context.Background(), &proto.String{Value: "client"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.GetValue())
}
