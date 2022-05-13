package distribute

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	client "go.etcd.io/etcd/client/v3"
)

var EtcdClient *client.Client

const requestTimeout = 10 * time.Second

func init() {
	if err := initClient(); err != nil {
		panic(err)
	}
}

func initClient() error {
	var etcdCfg = client.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	}
	var err error
	EtcdClient, err = client.New(etcdCfg)
	if err != nil {
		return err
	}
	return nil
}

func config() map[string]string {
	var cfg = map[string]string{
		"test_addr":             "127.0.0.1:1081",
		"test_aes_key":          "01B345B7A9ABC00F0123456789ABCXXXXAAAAAA",
		"test_https":            "0",
		"test_secret":           "",
		"test_private_key_path": "",
		"test_cert_file_path":   "",
	}
	return cfg
}

func getConfig(key string, isPrefix bool) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	var response *client.GetResponse
	var err error
	if isPrefix {
		response, err = EtcdClient.Get(ctx, key, client.WithPrefix())
	} else {
		response, err = EtcdClient.Get(ctx, key)
	}
	cancel()
	if err != nil {
		return nil, err
	}
	var config = make(map[string]string)
	for _, ev := range response.Kvs {
		config[string(ev.Key)] = string(ev.Value)
	}
	return config, nil
}

func putConfig(key string, val string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := EtcdClient.Put(ctx, key, val)
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			return fmt.Errorf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			return fmt.Errorf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			return fmt.Errorf("client-side error: %v", err)
		default:
			return fmt.Errorf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	return nil
}

func watchConfig(key string) {
	rch := EtcdClient.Watch(context.Background(), key, client.WithPrefix())
	for watcherRes := range rch {
		for _, ev := range watcherRes.Events {
			log.Printf("%s %q : %q", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
