package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// 过滤器
func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter:", info)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("painc: %v", r)
		}
	}()
	return handler(ctx, req)
}
