package main

import (
	"fmt"
	"math/rand"
	"net/rpc"
	"sync"
	"time"
)

const KVStoreServiceName string = "KVStoreService"

const Protocol string = "tcp"
const Ip string = "localhost"
const Port string = ":7777"

type KVStoreServiceInterface interface {
	Get(string, *string)
}

func RegisterKVStoreService() error {
	return rpc.RegisterName(KVStoreServiceName, NewKVStoreService())
}

type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m: make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	v, ok := p.m[key]
	if !ok {
		return fmt.Errorf("not found")
	}
	*value = v
	return nil
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]
	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

func (p *KVStoreService) Watch(timeout time.Duration, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	p.mu.Lock()
	p.filter[id] = func(key string) {
		ch <- key
	}
	p.mu.Unlock()

	var err error
	select {
	case <- time.After(timeout * time.Second):
		err = fmt.Errorf("timeout")
	case key := <- ch:
		*keyChanged = key
		err = nil
	}
	return err
}

