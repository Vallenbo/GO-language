package io

import (
	"fmt"
	"testing"
)

func Test_stdin(t *testing.T) {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		panic(err)
		return
	}

	print(s)
}
