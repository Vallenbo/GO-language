package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.5.3:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	_, err = cli.Put(ctx, "lmh", "lmh")
	cancel()
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case errors.Is(err, context.DeadlineExceeded):
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case errors.Is(err, rpctypes.ErrEmptyKey):
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "lmh")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
