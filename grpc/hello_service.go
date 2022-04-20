package main

import (
	"context"
	"fmt"
	"go-advanced/grpc/proto"
	"io"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 增加用户认证
type grpcServer struct {
	auth *Authentication
}

type HelloServiceImpl struct{}

func NewHelloServiceImpl() *grpcServer {
	return &grpcServer{
		auth: &Authentication{
			User:     "gopher",
			Password: "password",
		},
	}
}

func (p *grpcServer) Hello(ctx context.Context, args *proto.String) (*proto.String, error) {
	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}
	reply := &proto.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *grpcServer) Channel(stream proto.HelloService_ChannelServer) error {
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

type Authentication struct {
	User     string
	Password string
}

func (auth *Authentication) GetRequestMetadata(ctx context.Context, opts ...string) (map[string]string, error) {
	return map[string]string{"user": auth.User, "password": auth.Password}, nil
}

func (auth *Authentication) RequireTransportSecurity() bool {
	return false
}

func (auth *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var appid, appKey string
	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appKey = val[0]
	}

	if appid != auth.User || appKey != auth.Password {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *proto.StringMessage) (*proto.StringMessage, error) {
	return &proto.StringMessage{Value: "Get hi:" + message.GetValue() + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *proto.StringMessage) (*proto.StringMessage, error) {
	return &proto.StringMessage{Value: "Post hi:" + message.GetValue() + "@"}, nil
}
