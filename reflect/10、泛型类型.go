package main

import "fmt"

type List[T any] []T //泛型切片的定义

// 泛型集合的定义
type MapT[k comparable, v any] map[k]v //声明两个泛型，分别为k,v

type Chan[T any] chan T //泛型通道的定义

func main() {

}

type GetKey[T comparable] interface { //泛型接口
	any
	Get() T
}

func listToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(MapT[k, T], len(list))
	for _, data := range list {
		mp[data.Get()] = data
	}
	return mp
}

type user struct {
	ID   int64
	Name string
	Age  int
}

type address struct {
	ID       int
	Province string
	City     string
}

func (u user) Get() int64 {
	return u.ID
}

func (addr address) Get() int {
	return addr.ID

}

func InterfaceCase() {
	userList := []GetKey[int64]{
		user{ID: 1, Name: "nike", Age: 18},
		user{ID: 2, Name: "king", Age: 19},
	}
	addrList := []GetKey[int]{
		address{ID: 1, Province: "湖南", City: "湖南"},
		address{ID: 2, Province: "长沙", City: "长沙"},
	}
	userMp := listToMap[int64, GetKey[int64]](userList)
	fmt.Println(userMp)

	addrMp := listToMap[int, GetKey[int]](addrList)
	fmt.Println(addrMp)

}

type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

func (myStruct MyStruct[T]) GetData(t T) T {
	var i interface{} = 10
	a, ok := i.(int)
	//b, ok := t.(int) //泛型不支持断言
}