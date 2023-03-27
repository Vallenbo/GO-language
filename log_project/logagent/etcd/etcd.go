package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`  // 日志存放的路径
	Topic string `json:"topic"` // 日志发往kafka中的哪个Topic
}

func Init(addr string, timeout time.Duration) (err error) { // 初始化etcd的函数
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(addr, ";"),
		DialTimeout: timeout,
	})
	return
}

func GetConf(key string) (logEntryConf []*LogEntry, err error) { // 从etcd中获取日志收集项的配置信息
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%v\n", err)
			return
		}
	}
	return
}

func WatchConf(key string, newConfChan chan<- []*LogEntry) { // etcd watch
	rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
	for wresp := range rch {                    // 从通道尝试取值（监视的信息）
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)

			var newConf []*LogEntry //	通知taillog.taskMgr
			//1. 先判断操作的类型
			if ev.Type != clientv3.EventTypeDelete { // 如果不是是删除操作
				err := json.Unmarshal(ev.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed, err:%v\n", err)
					continue
				}
			}
			fmt.Printf("get new conf: %v\n", newConf)
			newConfChan <- newConf
		}
	}
}
