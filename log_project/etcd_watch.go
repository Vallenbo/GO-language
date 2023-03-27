package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func main() { // watch demo
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second, //超时时间
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// watch key:q1mi change  // 派一个哨兵 一直监视着 zhangyafei这个key的变化（新增 删除 修改））
	rch := cli.Watch(context.Background(), "zhangyafei") // 返回通道 <-chan WatchResponse

	for wresp := range rch { // 从通道尝试取值（监视的信息）
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
