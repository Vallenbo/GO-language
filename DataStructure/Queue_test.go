package DataStructure

import (
	"fmt"
	"testing"
)

type Queue struct { // 定义队列结构体
	data []interface{} // 数据
}

func (queue *Queue) Enqueue(data interface{}) { // 将元素加入队列
	queue.data = append(queue.data, data)
}

func (queue *Queue) Dequeue() interface{} { // 从队列中取出元素
	if len(queue.data) == 0 { // 如果队列为空，则返回nil
		return nil
	}

	data := queue.data[0]       // 取出队首元素
	queue.data = queue.data[1:] // 将队首元素从队列中删除
	return data
}

func (queue *Queue) Peek() interface{} { // 获取队首元素
	if len(queue.data) == 0 { // 如果队列为空，则返回nil
		return nil
	}

	data := queue.data[0] // 获取队首元素
	return data
}

func (queue *Queue) IsEmpty() bool { // 判断队列是否为空
	return len(queue.data) == 0
}
func Test_queue(*testing.T) {
	queue := &Queue{}
	queue.Enqueue(1) // 将元素加入队列
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Peek())    // 获取队首元素
	fmt.Println(queue.Dequeue()) // 从队列中取出元素
	fmt.Println(queue.IsEmpty()) // 判断队列是否为空
}
