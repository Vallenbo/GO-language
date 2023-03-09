package main

import "fmt"

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func switchDemo1() { //switch 条件判断
	finger := 3
	switch finger {
	case 1, 9: //多个条件用 ， 进行分隔
		fmt.Println("大拇指")
	case 'e': //字符用单引号，"asd" 字符串用双引号
		fmt.Println("食指")
		fallthrough //fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag //goto语句跳到 目标位置
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag: //goto XX  //XX: 程序块
	fmt.Println("结束for循环")
}

func main() {
	//if  != nil { break ;  continue
	//
	//} else if  {
	//
	//} else {}

	//for ;; { }

}
