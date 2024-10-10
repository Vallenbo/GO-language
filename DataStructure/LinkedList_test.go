package DataStructure

import (
	"fmt"
	"testing"
)

type Node struct { // 定义链表结构体
	data interface{} // 数据域
	next *Node       // 指针域
}

type LinkedList struct { // 定义链表类
	head *Node // 头指针
}

func (list *LinkedList) AddNode(data interface{}) { // 添加节点
	node := &Node{data: data}

	if list.head == nil { // 如果链表为空，则新节点为头节点
		list.head = node
	} else {
		last := list.head // 否则，从头节点开始遍历，找到最后一个节点
		for last.next != nil {
			last = last.next
		}
		last.next = node // 将新节点插入到最后一个节点的后面
	}
}

func (list *LinkedList) RemoveNode(data interface{}) { // 删除节点
	if list.head == nil { // 如果链表为空，则直接返回
		return
	}

	if list.head.data == data { // 如果要删除的是头节点，则将头指针指向下一个节点
		list.head = list.head.next
		return
	}

	prev := list.head // 从头节点开始遍历，找到要删除的节点的前一个节点
	for prev.next != nil {
		if prev.next.data == data {
			prev.next = prev.next.next // 将前一个节点的指针指向要删除的节点的下一个节点
			return
		}
		prev = prev.next
	}
}

func (list *LinkedList) Traverse() { // 遍历链表
	node := list.head // 从头节点开始遍历，直到遍历完整个链表
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}
func Test_LinkedList(*testing.T) {
	list := &LinkedList{}
	list.AddNode(1) // 添加节点
	list.AddNode(2)
	list.AddNode(3)

	list.Traverse()    // 遍历链表
	list.RemoveNode(2) // 删除节点
	list.Traverse()    // 再次遍历链表

}
