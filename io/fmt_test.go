package io

import (
	"fmt"
	"testing"
)

func Test_fmt(*testing.T) {
	var n = 100
	fmt.Printf("%T\n", n) // int
	fmt.Printf("%v\n", n) // 100
	fmt.Printf("%b\n", n) //1100100
	fmt.Printf("%d\n", n) //100
	fmt.Printf("%o\n", n) //144
	fmt.Printf("%x\n", n) //64

	var s = "hello 沙河！"
	fmt.Printf("字符串：%s\n", s)  //字符串：hello 沙河！
	fmt.Printf("字符串：%v\n", s)  //字符串：hello 沙河！
	fmt.Printf("字符串：%#v\n", s) //字符串："hello 沙河！"

}

func testPrint() { //1、Print打印函数

	fmt.Print("在终端打印该信息。")
	name := "沙河小王子"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}

func testScan() { //2、输入函数
	var (
		name    string
		age     int
		married bool
	)
	_, err := fmt.Scan(&name, &age, &married)                      //传入指针 //输入：小王子 28 false
	fmt.Printf("name:%s age:%d married:%t \n", name, age, married) //name:小王子 age:28 married:false

	_, err = fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)    //format指定格式 //1:小王子 2:28 3:false
	fmt.Printf("name:%s age:%d married:%t \n", name, age, married) //name:小王子 age:28 married:false

	_, err = fmt.Scanln(&name, &age, &married)
	if err != nil {
		return
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	if err != nil {
		return
	}
}
