package main

import (
	"fmt"
	"strconv"
)

var (
	name string
	age  int
	isOK bool
)

// 类型推导
var name1 = "李明"

// 常量
const name2 = "小白"
const name3 = "xxx"
const (
	A = iota //出现时为0
	B
	C
	D
)

const (
	a, b = iota, iota + 1 //每新增一行才加1
	c, d = iota + 2, iota + 3
)

const ( //定义数量级
	_  = iota             //出现的时候为0,赋值隐藏，扔掉了
	KB = 1 << (10 * iota) //iota=1、基础内容 ，该行表示1左移十位：10000000000
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	for index, value := range name2 {
		fmt.Printf("index: %v ,stringvalue :%v ,value :%T\n", index, string(value), value)
	}
	fmt.Printf("%v\n", name2+name3)
	fmt.Printf("%v\n", strconv.Itoa(1))
}

/* 输出：
index: 0 ,stringvalue :小 ,value :int32
index: 3 ,stringvalue :白 ,value :int32
小白xxx
*/
