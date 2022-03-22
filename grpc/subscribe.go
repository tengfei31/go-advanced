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
)

func main() {
	//subscribeBaseTLS()
	subscribeClientTLS()
}

func subscribeClientTLS() {
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

	subscribe(conn)
}

func subscribeBaseTLS() {
	creds, err := credentials.NewClientTLSFromFile("secret/server.crt", "server.grpc.io")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	subscribe(conn)
}

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
