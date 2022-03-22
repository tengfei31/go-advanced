package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"go-advanced/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	//mainClientTLS()

	//mainServerTLS()

	//mainNoTLS()

	mainBaseGrpc()
}

// 利用客户端证书和CA根证书、服务器名字验证
func mainClientTLS() {
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
func mainServerTLS() {
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
func mainNoTLS() {
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
func mainBaseGrpc() {
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
