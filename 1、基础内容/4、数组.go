package main

import "fmt"

func main() {
	//数组初始化方法一
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1、基础内容 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]

	//数组初始化方法二
	var testArray1 [3]int
	var numArray1 = [...]int{1, 2}
	var cityArray1 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray1)                          //[0 0 0]
	fmt.Println(numArray1)                           //[1、基础内容 2]
	fmt.Printf("type of numArray:%T\n", numArray1)   //type of numArray:[2]int
	fmt.Println(cityArray1)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray1) //type of cityArray:[3]string

	//数组初始化方法三
	a1 := [...]int{1: 1, 3: 5}
	fmt.Println(a1)                  // [0 1、基础内容 0 5]
	fmt.Printf("type of a:%T\n", a1) //type of a:[4]in

	//数组的遍历，方法1：通过索引
	var a2 = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}

	// 方法2：for range遍历
	for index, value := range a2 {
		fmt.Println(index, value)
	}
	//数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
	//[n]*T表示指针数组，*[n]T表示数组指针 。
}

func Mutarrary() { //多维数组
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆

	for _, v1 := range a { //多维数组的遍历
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	////多维数组支持的写法
	//a := [...][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	////不支持多维数组的内层使用...
	//b := [3][...]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
}
