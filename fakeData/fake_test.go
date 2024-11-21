package fakeData

import (
	"fmt"
	"github.com/icrowley/fake"
	"testing"
)

func Test_fakeIP4(t *testing.T) {
	for i := 0; i < 10; i++ {
		ip := fake.IPv4()
		fmt.Println("随机IPv4地址:", ip)
	}
}
