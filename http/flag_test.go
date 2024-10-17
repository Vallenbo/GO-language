package http

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_Arges(t *testing.T) { // os.Args 命令行获取参数
	fmt.Printf("%#v\n\n", os.Args)     // []string{"3、flag包--命令行参数.exe"}
	fmt.Printf(os.Args[0], os.Args[1]) //获取第一个参数，和第二个参数 ：___go_build_42.exe ，a ,b
	fmt.Printf("%T", os.Args)          //[]string
	flag.Parse()                       //解析命令行参数
	fmt.Println(flag.Args())           //返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())           //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag())          //返回使用的命令行参数个数
}

func typeArgs() { // 2- 设置命令行参数默认值
	name := flag.String("name", "张三", "姓名") // -name:命令行参数 , value:默认值 ，usage:提示内容  //返回指针
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")
	flag.Parse() //解析从操作系统解析命令行标志
	fmt.Println(*name, *age, *married, *delay)
	// $ ./args_demo -name “李白” //指定参数
}

func vartypeArgs() { //2-2 设置命令行参数，以变量确定命令行类型
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")
	flag.Parse() //解析从操作系统解析命令行标志
	fmt.Println(name, age, married, delay)
	// $ ./args_demo -name “李白” //指定参数
}

//命令行参数使用提示：
//
//$ ./flag_demo -help
//Usage of ./flag_demo:
//  -age int
//        年龄 (default 18)
//  -d duration
//        时间间隔
//  -married
//        婚否
//  -name string
//        姓名 (default "张三")
//正常使用命令行flag参数：
//
//$ ./flag_demo -name 沙河娜扎 --age 28 -married=false -d=1h30m
//沙河娜扎 28 false 1h30m0s
//[]
//0
//4
//使用非flag命令行参数：
//
//$ ./flag_demo a b c
//张三 18 false 0s
//[a b c]
//3
//0
