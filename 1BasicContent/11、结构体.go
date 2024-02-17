package main

import "fmt"

type MyInt int // 将MyInt定义为int类型
//type int32 rune //官方内置别名

type person struct { // person称为类型 struct称为结构体
	name  string
	city  string
	age   int8
	hobby []string
}

type person1 struct {
	name, city string
	age        int8
}

func f(x person) { //
	x.age = 30 //函数收到的值，永远是cp复制的值   //修改需要用指针
}

func f10(x *person) { //3-结构体用指针传参
	(*x).age = 30 //方法1：修改指针内的数据
	x.age = 30    //方法2：go语言的语法糖，自动判断接受类型是否为指针
}

func newPerson(name, city string, age int8) *person { //1-2结构体初始化，使用构造函数返回值（使用指针类型）
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	var p1 person //1-1结构体初始化,使用键值对
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18

	user := struct { //1-2使用匿名结构体，并进行赋值
		Name string
		Age  int
	}{"小王子", 18}

	p5 := person1{"x", "hunan", 15} //1-3使用匿名结构体，并进行赋值

	p6 := newPerson("nazha", "guandong", 25)

	var u = person1{"x", "x", 25} //2-3使用结构体直接赋值

	fmt.Printf("p1=%T,u:%T ,p5:%T,p6: %T\n", p1, u, p5, p6) //打印类型：p1=main.person
	fmt.Printf("p1=%v\n", p1)                               //p1={沙河娜扎 北京 18}
	fmt.Printf("p1=%#v\n", p1)                              //p1=main.person{name:"沙河娜扎", city:"北京", age:18}

	var p2 = new(person) //3-1创建指针类型结构体
	p3 := &person{       //3-2对结构体取地址
		name: "小王子",
		city: "北京",
		age:  18}

	fmt.Printf("%#v\n", user)
	fmt.Printf("p2：%T，p3:%T\n", p2, p3) //*main.person
	fmt.Printf("p3=%#v\n", p3)          //p3=&main.person{name:"", city:"", age:0}

}
