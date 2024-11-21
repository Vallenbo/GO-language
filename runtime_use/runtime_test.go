package runtime_use

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func Test_runtime(t *testing.T) {
	myFunction(42, "hello")
}

func myFunction(arg1 int, arg2 string) {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(funcName, ".")
	functionName := parts[len(parts)-1]
	fmt.Printf("当前函数名：%s\n", functionName)

	// 获取参数
	fmt.Printf("传入参数：arg1=%d, arg2=%s\n", arg1, arg2)
}
