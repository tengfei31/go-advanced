package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func RegisterRpc() {
	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalln("listen fail", err)
	}
	log.Println("listen success tcp://0.0.0.0:6666")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("accept fail", err)
		}
		log.Println("accept:", conn.RemoteAddr().String())
		//go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}

func RegisterHttpRpc() {
	RegisterHelloService(new(HelloService))
	http.HandleFunc("/jsonrpc", func(response http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     response,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":6666", nil)
}

func RegisterKVStoreRpc() {
	RegisterKVStoreService()
	listener, err := net.Listen(Protocol, Ip+Port)
	if err != nil {
		log.Fatalln("listen fail", err)
	}
	log.Printf("listen success %s://%s", Protocol, listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("accept fail", err)
		}
		log.Println("accept:", conn.RemoteAddr().String())
		go rpc.ServeConn(conn)
	}
}

func reverseProxy() {
	rpc.Register(new(HelloService))

	for {
		conn, _ := net.Dial("tcp", ":1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		log.Println("connect:", conn.RemoteAddr())
		rpc.ServeConn(conn)
		log.Println("connect success")
		conn.Close()
	}
}

func RegisterContextRpc() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			p.Register(&HelloService{conn: conn})
			p.ServeConn(conn)
		}()
	}
}
