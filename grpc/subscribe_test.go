package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TestSubscribeClientTLS(t *testing.T) {
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

func TestSubscribeBaseTLS(t *testing.T) {
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

func TestSubscribeNoTLS(t *testing.T) {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	subscribe(conn)
}
