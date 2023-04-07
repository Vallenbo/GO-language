package main

import (
	"fmt"
	"math"
)

var a1 int = 10
var b1 int = 077
var c1 int = 0xff
var bool1 bool //布尔型，只能进行判断

func main() {

	fmt.Printf("八进制数%d\n", b1)  //63
	fmt.Printf("十进制数%d\n", a1)  //10
	fmt.Printf("十六进制数%d\n", c1) //255

	//浮点型
	fmt.Printf("%f\n", math.Pi)   //3.141593
	fmt.Printf("%.2f\n", math.Pi) //3.14

	//字符串 , 1字节 =8bit
	// s3 := 'w'  // 字符用‘’
	s1 := "hello"
	s3 := "你好"
	fmt.Printf("%s，%s\n", s1, s3) //hello，你好

	//字符串转义字符
	//\r	回车符（返回行首）
	//\n	换行符（直接跳到下一行的同列位置）
	//\t	制表符
	//\'	单引号
	//\"	双引号
	//\\	反斜杠
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")   //str := "c:\Code\lesson1\go.exe"
	fmt.Println("'str := \"c:\\Code\\lesson1\\go.exe\"'") //'str := "c:\Code\lesson1\go.exe"'

	// ``反引号
	s2 := ` 
	事态炎凉
	`
	println(s2)
	ss := s1 + s3   //字符串相加
	fmt.Println(ss) //hello你好

	for _, s01 := range ss { //打印单个字符
		fmt.Printf("%c ", s01) //hello你好
	}

	// #strings.Contains(str1, "s") //判断字符串中 是否有指定字符
}
