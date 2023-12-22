package main

import "fmt"

func test() { //2、切片常用函数
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]                                                       // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))       //	s:[2 3] len(s):2 cap(s):4
	s2 := s[3:4]                                                      // 注意：索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2)) //s2:[5] len(s2):1、基础内容 cap(s2):1、基础内容
}

func main() {
	// 1、基础内容、声明切片类型,初始化
	var a []string              //声明一个字符串切片
	var c = []bool{false, true} //声明一个布尔切片并初始化
	c[0] = false                //初始化
	c[1] = true
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true ，说明未开辟物理地址
	fmt.Println(c == nil) //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	// 3、切片完整表达式
	a1 := [5]int{1, 2, 3, 4, 5}
	t := a1[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t)) //t:[2 3] len(t):2 cap(t):4

	//4、使用make函数创造切片
	s1 := make([]int, 5, 5)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d", s1, len(s1), cap(s1)) //s1=[0 0 0 0 0] len(s1)=5 cap(s1)=5

	//5、未分配地址的切片，都是为nil
	//var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	//s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	//s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil

}

func Advanced_operation() { //6、切片进阶操作，追加append
	var s []int
	s = append(s, 1)       // [1、基础内容]
	s = append(s, 2, 3, 4) // [1、基础内容 2 3 4]
	s2 := []int{5, 6, 7}
	s = append(s, s2...) // [1、基础内容 2 3 4 5 6 7]

	var citySlice []string                          //切片拥有自动扩容策略
	citySlice = append(citySlice, "北京")             // 追加一个元素
	citySlice = append(citySlice, "上海", "广州", "深圳") // 追加多个元素
	a := []string{"成都", "重庆"}                       // 追加切片
	citySlice = append(citySlice, a...)             //...表示将a切片拆开进行添加
	fmt.Println(citySlice)                          //[北京 上海 广州 深圳 成都 重庆]
}

func test1() { //6-2、copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1、基础内容 2 3 4 5]
	fmt.Println(c) //[1、基础内容 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1、基础内容 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}

func test2() { //6-3、删除切片中元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37} // 从切片中删除元素
	a = append(a[:2], a[3:]...)                // 要删除索引为2的元素
	fmt.Println(a)                             //[30 31 33 34 35 36 37]
}
