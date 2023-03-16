package main

import "fmt"

type Person struct { //Person 结构体，首字母大写为公有，外部包可见
	name string
	age  int8
}

type Person1 struct { //Person 结构体Person类型
	string //结构体匿名字段，即没有字段名
	int
}

func NewPerson(name string, age int8) *Person { //NewPerson 构造函数
	return &Person{
		name: name,
		age:  age,
	}
}

//func (p Person) SetAge2(newAge int8) { // 使用值接收者
//	p.age++ //改变的是cp副本
//}

func (p *Person) Dream() { //2-使用指针类型，Dream Person做梦的方法,
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
	p.age++ //实现原数据更改
}

type Address struct { // Address 地址结构体
	Province string
	City     string
}

type User struct { // 3-结构体嵌套，User 用户结构体
	Name    string
	Gender  string
	Address Address
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()

	//p1 := Person1{  //2-匿名字段结构体赋值
	//	"小王子",
	//	18,
	//}
	//fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
	//fmt.Println(p1.string, p1.int) //北京 18

	user1 := User{ //3-结构体嵌套赋值
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}
	fmt.Printf("user1=%#v\n", user1) //user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

	var user2 User1 // #2-2结构体嵌套匿名字段，进行赋值
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
}

type Address1 struct { // Address 地址结构体
	Province string
	City     string
}

type User1 struct { // # 2-2结构体嵌套匿名字段，User 用户结构体
	Name    string
	Gender  string
	Address //匿名字段
}
