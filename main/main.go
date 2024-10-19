package main

import (
	"fmt"
)

func main() {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		panic(err)
		return
	}

	print(s)

}
