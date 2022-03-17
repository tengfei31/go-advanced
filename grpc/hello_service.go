package main

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	"io"
	"strings"
	"time"
)
import "go-advanced/grpc/proto"

type HelloServiceImpl struct{}

func NewHelloServiceImpl() *HelloServiceImpl {
	return &HelloServiceImpl{}
}

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

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(ctx context.Context, arg *proto.String) (*proto.String, error) {
	p.pub.Publish(arg.GetValue())
	return &proto.String{}, nil
}

func (p *PubsubService) Subscribe(arg *proto.String, stream proto.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for value := range ch {
		err := stream.Send(&proto.String{Value: value.(string)})
		if err != nil {
			return err
		}
	}

	return nil
}
