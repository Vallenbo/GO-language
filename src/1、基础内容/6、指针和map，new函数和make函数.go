package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type类型:*int
	fmt.Println(&b)                    // 0xc00000e018

	//1、基础内容、进阶操作，创造指针，使用new和make和map函数
	var a1 *int //定义
	a1 = new(int)
	*a1 = 100        //赋值
	fmt.Println(*a1) //为nil,因为没有开辟物理地址，需要使用new函数

	var b2 map[string]int
	b2 = make(map[string]int, 10) //内存分配
	//b2 := make(map[string]int, 10) //一次性赋值
	b2["沙河娜扎"] = 100
	fmt.Println(b2) //map[沙河娜扎:100]

	//Map判断某个键值对是否存在
	v, ok := b2["张三"] //ok为bool类型
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人") //查无此人
	}
	delete(b2, "小明") //将 小明:100 从map中删除
}

func test3() { //Map的遍历
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

func Test_advont() { //map和slice结合
	s1 := make([]map[int]string, 2) //元素类型为map的切片
	//需要分别先对1、slice进行初始化，再对2、map进行初始化
	s1[0] = make(map[int]string, 1)
	s1[0][0] = "A" //添加一个键值对
	fmt.Print(s1)

	//元素类型为切片的map
	s2 := make(map[string][]int, 5)
	s2["A"] = []int{1, 2}
	fmt.Print(s2)
}
