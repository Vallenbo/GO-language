package io

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_strconv(*testing.T) {
	s1 := "100"
	i1, err := strconv.Atoi(s1) //字符串类型的整数转换为int类型
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	i2 := 200
	s2 := strconv.Itoa(i2)                    //int类型数据转换为对应的字符串
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
}

func ParseSeries() { //Parse系列函数
	//b, err := strconv.ParseBool("true")
	//f, err := strconv.ParseFloat("3.1415", 64)
	//i, err := strconv.ParseInt("-2", 10, 64)
	//u, err := strconv.ParseUint("2", 10, 64)
}

func FormatSeries() { //Format系列函数
	//s1 := strconv.FormatBool(true)
	//s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	//s3 := strconv.FormatInt(-2, 16)
	//s4 := strconv.FormatUint(2, 16)
}

func Other() {

}
