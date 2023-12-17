package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {

	//dial, err := rpc.Dial("tcp", "192.168.0.5:8800")
	dial, err := jsonrpc.Dial("tcp", "192.168.0.5:8800")
	if err != nil {
		fmt.Println("the error of rpc Dial is :", err)
		return
	}

	var reply string
	err = dial.Call("Hello.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("the error of dial Call is :", err)
		return
	}

	fmt.Printf(reply)
}