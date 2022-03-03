package main

import (
	"context"
	"io"
)
import "go-advanced/grpc/proto"

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *proto.String) (*proto.String, error) {
	reply := &proto.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream proto.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &proto.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
