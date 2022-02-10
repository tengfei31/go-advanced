package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName string = "HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

func (hs *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	hs.isLogin = true
	return nil
}

// Hello 增加hello前缀
func (hs *HelloService) Hello(request string, reply *string) error {
	if !hs.isLogin {
		return fmt.Errorf("plase login")
	}
	*reply = "hello:" + request + ", from:" + hs.conn.RemoteAddr().String()
	return nil
}

// Hello 增加hello前缀
//func (hs *HelloService) Hello(request *proto.String, reply *proto.String) error {
//	reply.Value = "hello:" + request.GetValue()
//	return nil
//}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Login(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Login", request, reply)
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
