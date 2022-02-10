package main

import "context"
import "go-advanced/grpc/proto"

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *proto.String) (*proto.String, error) {
	reply := &proto.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}
