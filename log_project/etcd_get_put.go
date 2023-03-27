package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func GetOurboundIP() (ip string, err error) { // 获取本地对外IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}

func main() { // etcd client put/get demo
	cli, err := clientv3.New(clientv3.Config{ // use etcd/client/v3
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second, //超时时间
	})
	if err != nil { // handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second) //创建一个释放资源的context

	ip, err := GetOurboundIP() //初始化key
	if err != nil {
		panic(err)
	}
	logConfKey := fmt.Sprintf("/logagent/%s/collect_config", ip)

	//配置信息
	value := `[{"path":"F:/go/nginx.log","topic":"web_log"},{"path":"F:/go/redis.log","topic":"redis_log"},{"path":"F:/go/mysql.log","topic":"mysql_log"}]`
	//value := `[{"path":"f:/tmp/nginx.log","topic":"web_log"},{"path":"f:/tmp/redis.log","topic":"redis_log"}]`
	//_, err = cli.Put(ctx, "zhangyafei", "dsb")
	_, err = cli.Put(ctx, logConfKey, value) // put

	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second) // get

	resp, err := cli.Get(ctx, logConfKey)
	//resp, err := cli.Get(ctx, "/logagent/collect_config")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
