package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 专门往kafka写日志的模块

type LogData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer // 声明一个全局的连接kafka的生产者client
	logDataChan chan *LogData
)

func Init(addrs []string, chanMaxSize int) (err error) { // init初始化client
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	client, err = sarama.NewSyncProducer(addrs, config) // 连接kafka
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	logDataChan = make(chan *LogData, chanMaxSize) // 初始化logDataChan
	go SendToKafka()                               // 开启后台的goroutine，从通道中取数据发往kafka
	return
}

func SendToChan(topic, data string) { // 给外部暴露的一个函数，噶函数只把日志数据发送到一个内部的channel中
	msg := &LogData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

func SendToKafka() { // 真正往kafka发送日志的函数
	for {
		select {
		case logData := <-logDataChan:
			msg := &sarama.ProducerMessage{} // 构造一个消息
			msg.Topic = logData.topic
			msg.Value = sarama.StringEncoder(logData.data)

			pid, offset, err := client.SendMessage(msg) // 发送到kafka
			if err != nil {
				fmt.Println("sned msg failed, err:", err)
			}
			fmt.Printf("send msg success, pid:%v offset:%v\n", pid, offset) //fmt.Println("发送成功")
		}
	}

}
