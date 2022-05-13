package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func Client() {
	client, err := DialHelloService("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln("dialing:", err)
	}
	var loginReply string
	err = client.Login("user:password", &loginReply)
	if err != nil {
		log.Fatalln("call HelloService.Login fail:", err)
	}
	var reply string
	err = client.Hello("client", &reply)
	if err != nil {
		log.Fatalln("call HelloService.Hello fail:", err)
	}
	log.Println("client recv:", reply)
}

func JsonClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Fatalln("dialing:", err)
	}
	var client = rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err = client.Call(HelloServiceName+".Hello", "client", &reply)
	if err != nil {
		log.Fatalln("call HelloService.Hello fail:", err)
	}
	log.Println("client recv:", reply)
}

func kvStoreClient() {
	conn, err := net.Dial(Protocol, Ip+Port)
	if err != nil {
		log.Fatalln("dialing:", err)
	}
	var client = rpc.NewClient(conn)
	var value string
	err = client.Call(KVStoreServiceName+".Get", "abc", &value)
	if err != nil {
		log.Fatalf("call %s.Get fail:%v", KVStoreServiceName, err)
	}
	doClientWork(client)
	log.Println("client recv:", value)
}

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call(KVStoreServiceName+".Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("watch:", keyChanged)
	}()
	err := client.Call(KVStoreServiceName+".Set", [2]string{"abc", "abc-value"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)
}

func proxyClient() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp error:", err)
	}
	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("accept error:", err)
			}

			clientChan <- rpc.NewClient(conn)
		}
	}()

	doClientWorkProxy(clientChan)
}

func doClientWorkProxy(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()

	var reply string
	err := client.Call(HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}
